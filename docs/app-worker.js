const cacheName = "app-" + "29c4bec5274f0465c3f81ab51458fbb946cec4ef";

self.addEventListener("install", event => {
  console.log("installing app worker 29c4bec5274f0465c3f81ab51458fbb946cec4ef");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/",
          "/app.css",
          "/app.js",
          "/manifest.webmanifest",
          "/wasm_exec.js",
          "/web/app.wasm",
          "/web/main.css",
          "https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css",
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
  console.log("app worker 29c4bec5274f0465c3f81ab51458fbb946cec4ef is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
