// https://coderun.yandex.ru/selections/first-2023-frontend/problems/extend-transport-system
// ExtendTransportSystem - problem 35
function extendTransportSystem(EarthRoute, MoonRoute) {
    const arr = [];
  
      const handler = {
          set: function(obj, prop, value){
              if(value !== obj.length) {
                  arr.push({...value, origin: value.destination, destination: 'Mothership'});
              }
              obj[prop] = value;
              return true;
          }
      };
  
      EarthRoute.vault = new Proxy(EarthRoute.vault, handler);
      MoonRoute.warehouse = new Proxy(MoonRoute.warehouse, handler);
  
      return arr;
  }
  
  module.exports = extendTransportSystem
  