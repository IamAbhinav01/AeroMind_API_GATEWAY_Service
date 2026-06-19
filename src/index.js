const express = require('express');
const { ServerConfig, LoggerConfig } = require('./config');
const apiRoutes = require('./routes');
const cors = require('cors');
const { FRONTEND_SERVICE_URL } = require('./config/server.config');
const app = express();

app.use(
  cors({
    origin: FRONTEND_SERVICE_URL,
    methods: ['GET', 'POST', 'PUT', 'PATCH', 'DELETE'],
    allowedHeaders: ['Authorization', 'Content-Type'],
  })
);

app.use('/api/v1/user', express.json());
app.use('/api/v1/user', express.urlencoded({ extended: true }));
app.use('/api', apiRoutes);
app.listen(ServerConfig.PORT, () => {
  console.log(`server started at port: ${ServerConfig.PORT}`);
  LoggerConfig.info(`server started at port: ${ServerConfig.PORT}`);
});
