console.log("newwewe")

window.addEventListener('beforeinstallprompt', e => {
    console.log('beforeinstallprompt Event fired');
    e.preventDefault();
    // Stash the event so it can be triggered later.
    this.deferredPrompt = e;
    return false;
});
