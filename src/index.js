const express = require('express');
const path    = require('path');
const { ServerConfig, LoggerConfig } = require('./config');
const apiRoutes = require('./routes');
const cors = require('cors');
const { FRONTEND_SERVICE_URL } = require('./config/server.config');
const app = express();

app.use(
  cors({
    origin: (origin, callback) => {
      // Allow requests from the configured frontend URL, VS Code Live Server,
      // file:// protocol (origin = null), or any localhost dev server,
      // and localhost:4005 itself (when the gateway serves the static files)
      const allowed = [
        FRONTEND_SERVICE_URL,
        `http://localhost:${ServerConfig.PORT}`,
        `http://127.0.0.1:${ServerConfig.PORT}`,
        'http://localhost:5500',
        'http://127.0.0.1:5500',
        'http://localhost:3001',
        'null',
      ].filter(Boolean);
      if (!origin || allowed.includes(origin)) return callback(null, true);
      return callback(new Error(`CORS: origin ${origin} not allowed`));
    },
    methods: ['GET', 'POST', 'PUT', 'PATCH', 'DELETE'],
    allowedHeaders: ['Authorization', 'Content-Type', 'x-access-token', 'x-idempotency-key'],
  })
);

// ── Serve the FRONTEND as static files ─────────────────────────────────────
// All .html, .css, .js and assets (aeromind.mp4) in AeroMind_FRONTEND are
// served directly via the Gateway.  Navigate to http://localhost:4005 to open
// the app — no separate Live Server needed.
const FRONTEND_DIR = path.join(__dirname, '..', '..', 'AeroMind_FRONTEND');
app.use(express.static(FRONTEND_DIR));

app.use('/api/v1/user', express.json());
app.use('/api/v1/user', express.urlencoded({ extended: true }));
app.use('/api', apiRoutes);

// ── Fallback: serve index.html for any unknown non-API route ───────────────
app.use((req, res, next) => {
  if (req.path.startsWith('/api')) return next();
  res.sendFile(path.join(FRONTEND_DIR, 'index.html'));
});

app.listen(ServerConfig.PORT, () => {
  LoggerConfig.info(`Gateway running at http://localhost:${ServerConfig.PORT}`);
  LoggerConfig.info(`Frontend static dir: ${FRONTEND_DIR}`);
  LoggerConfig.info(`Open app at:  http://localhost:${ServerConfig.PORT}/index.html`);
});

