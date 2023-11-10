declare var self: ServiceWorkerGlobalScope;
export {};

import CachedAssetsCatalog from 'asset/CachedAssetsCatalog';

function isNavigationRequest(request: Request): boolean {
    const acceptHeader = request.headers.get('Accept');
    return null !== acceptHeader && acceptHeader.includes('text/html');
}

function isStaticAssetRequest(request: Request): boolean {
    return request.method === 'GET' && (new URL(request.url)).pathname.startsWith('/compiled/');
}

async function upgradeRequest(cachedAssetsCatalog: CachedAssetsCatalog, request: Request): Promise<Response> {
    const headers = new Headers(request.headers);

    const assetDigest = await cachedAssetsCatalog.getAll();

    headers.set('X-Asset-Digest', JSON.stringify(assetDigest));

    const upgradedRequest = new Request(request, {
        headers: headers,
        mode: 'same-origin',
    });
    return fetch(upgradedRequest);
}

self.addEventListener('fetch', async (event: FetchEvent) => {
    const request = event.request;
    const cachedAssetsCatalog = new CachedAssetsCatalog();

    if (isNavigationRequest(request)) {
        event.respondWith(upgradeRequest(cachedAssetsCatalog, request));
        return;
    }

    if (isStaticAssetRequest(request)) {
        event.respondWith(caches.match(request).then(response => response || fetch(request)));
    }
});
