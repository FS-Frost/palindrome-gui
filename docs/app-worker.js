const cacheName = "app-" + "55c9622d9ff1839a8ce956875bb93170553afda3";

self.addEventListener("install", event => {
  console.log("installing app worker 55c9622d9ff1839a8ce956875bb93170553afda3");

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
          "/palindrome-gui/web/main.css",
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
  console.log("app worker 55c9622d9ff1839a8ce956875bb93170553afda3 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
