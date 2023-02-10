var Bugsnag = require('@bugsnag/js')
var BugsnagPluginExpress = require('@bugsnag/plugin-express')


var createError = require('http-errors');
var express = require('express');
var path = require('path');
var cookieParser = require('cookie-parser');
var logger = require('morgan');

const { expressjwt: expressJwt } = require('express-jwt');
const jwks = require('jwks-rsa');

var indexRouter = require('./routes/index');
var usersRouter = require('./routes/users');

var app = express();

const port = process.env.PORT || 50000;

Bugsnag.start({
  apiKey: '2ba623624a33a894a56ece104eb75f75',
  plugins: [BugsnagPluginExpress]
})
var middleware = Bugsnag.getPlugin('express')
app.use(middleware.requestHandler)


// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'jade');

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

var jwtCheck = expressJwt({
  secret: jwks.expressJwtSecret({
    cache : true,
    ratelimit : true,
    jwksRequestsPerMinute : 5,
    jwksUri : "https://dev-aa5pftowvhf5q3p2.us.auth0.com/.well-known/jwks.json"
  }),
  audience : "muscleapp",
  issuer : "https://dev-aa5pftowvhf5q3p2.us.auth0.com/",
  algorithms : ['RS256']
})

app.use(jwtCheck);
app.use('/', indexRouter);
app.use('/users', usersRouter);

// catch 404 and forward to error handler
app.use(function(req, res, next) {
  next(createError(404));
});

// error handler
app.use(function(err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.render('error');
});
app.listen(port);
app.use(middleware.errorHandler)

module.exports = app;
