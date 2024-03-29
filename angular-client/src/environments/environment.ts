// This file can be replaced during build by using the `fileReplacements` array.
// `ng build --prod` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.

export const environment = {
  production: false,
  usersApi: 'http://localhost:8000/api',
  adsApi: 'http://localhost:8001/api/ads',
  eventsApi: 'http://localhost:8002/api/events',
  commentsApi: 'http://localhost:8003/api/comments',

  loginRoute: 'login',
  registerRoute: 'register',
  profileRoute: 'profile',
  adListRoute: 'ad-list',
  adFormRoute: 'ad-form',
  adPageRoute: 'ad-page',
  eventFormRoute: 'event-form'
};

/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/dist/zone-error';  // Included with Angular CLI.
