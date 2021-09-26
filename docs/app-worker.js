const cacheName = "app-" + "8d1a18ecd89aefe48201ce0d18722546368d6b5b";

self.addEventListener("install", event => {
  console.log("installing app worker 8d1a18ecd89aefe48201ce0d18722546368d6b5b");
  self.skipWaiting();

  event.waitUntil(
    caches.open(cacheName).then(cache => {
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
  console.log("app worker 8d1a18ecd89aefe48201ce0d18722546368d6b5b is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
