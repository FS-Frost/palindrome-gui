const cacheName = "app-" + "e07af7d30589ca4903ae1a81799c92c6c321fa23";

self.addEventListener("install", event => {
  console.log("installing app worker e07af7d30589ca4903ae1a81799c92c6c321fa23");

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
  console.log("app worker e07af7d30589ca4903ae1a81799c92c6c321fa23 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
