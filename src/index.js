const express = require('express');
const { ServerConfig, LoggerConfig } = require('./config');
const apiRoutes = require('./routes');
const app = express();

// Body parsers are intentionally NOT applied globally.
// Applying them globally would consume the request stream before the proxy
// can forward it to downstream services (flights/bookings).
// They are applied only on the /user routes where the gateway itself handles the body.
app.use('/api/v1/user', express.json());
app.use('/api/v1/user', express.urlencoded({ extended: true }));
app.use('/api', apiRoutes);
app.listen(ServerConfig.PORT, () => {
  console.log(`server started at port: ${ServerConfig.PORT}`);
  LoggerConfig.info(`server started at port: ${ServerConfig.PORT}`);
});
