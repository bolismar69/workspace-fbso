declare const angular: any;

angular
  .module("cnpjApp", [])
  .directive("cnpjMask", function () {
    return {
      restrict: "A",
      require: "ngModel",
      link: function (_scope: any, element: any, _attrs: any, ngModel: any) {
        function applyMask(viewValue: any) {
          var masked = formatCnpjMasked(viewValue);
          if (masked !== viewValue) {
            ngModel.$setViewValue(masked);
            ngModel.$render();
          }
          return masked;
        }

        ngModel.$parsers.push(applyMask);
        ngModel.$formatters.push(function (modelValue: any) {
          return formatCnpjMasked(modelValue);
        });

        element.on("blur", function () {
          var masked = formatCnpjMasked(ngModel.$viewValue);
          if (masked !== ngModel.$viewValue) {
            _scope.$apply(function () {
              ngModel.$setViewValue(masked);
              ngModel.$render();
            });
          }
        });
      },
    };
  })
  .controller("MainController", function (this: any) {
    var vm: any = this;

    vm.cnpj = "";
    vm.normalized = null as string | null;
    vm.valid = false;

    vm.onChange = function () {
      vm.normalized = normalizeCnpj(vm.cnpj);
      vm.valid = isValidCnpj(vm.cnpj);
    };

    vm.inputOrNull = function () {
      return vm.cnpj === "" ? null : vm.cnpj;
    };

    vm.normalizedOrNull = function () {
      return vm.normalized;
    };

    vm.onChange();
  });
