const cacheName = "app-" + "cb52211d67fd2c7f94d0fd3916ba0c7a4e0aaf19";

self.addEventListener("install", event => {
  console.log("installing app worker cb52211d67fd2c7f94d0fd3916ba0c7a4e0aaf19");

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
  console.log("app worker cb52211d67fd2c7f94d0fd3916ba0c7a4e0aaf19 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
