'use strict';

/**
 * @ngdoc function
 * @name publicApp.controller:AboutCtrl
 * @description
 * # AboutCtrl
 * Controller of the publicApp
 */
angular.module('publicApp')
  .controller('AboutCtrl', function ($scope, Restangular) {
    var user = Restangular.one('users', 0);
    $scope.users = [user.get().$object];
  });
