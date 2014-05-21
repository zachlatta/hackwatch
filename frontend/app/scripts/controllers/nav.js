'use strict';

angular.module('frontendApp')
  .controller('NavCtrl', function ($scope, $location) {
    $scope.isActive = function (viewLocation) {
      return viewLocation === $location.path();
    };
  });
