const cacheName = "app-" + "503cf38b6501d81eccf4a78027b5751675a3714c";

self.addEventListener("install", event => {
  console.log("installing app worker 503cf38b6501d81eccf4a78027b5751675a3714c");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/palindrome-gui",
          "/palindrome-gui/app.css",
          "/palindrome-gui/app.js",
          "/palindrome-gui/manifest.webmanifest",
          "/palindrome-gui/wasm_exec.js",
          "/palindrome-gui/web/app.wasm",
          "https://storage.googleapis.com/murlok-github/icon-192.png",
          "https://storage.googleapis.com/murlok-github/icon-512.png",
          
        ]);
      }).
      then(() => {
        self.skipWaiting();
      })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker 503cf38b6501d81eccf4a78027b5751675a3714c is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
