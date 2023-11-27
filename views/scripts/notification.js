const urlParams = new URLSearchParams(window.location.search);
const notifStatus = urlParams.get('status');
const reason = urlParams.get('reason');

const timeout = 5000

if (notifStatus) {
    if (notifStatus == "success") {
        successNotif.classList.replace('hidden', 'flex')
        setTimeout(() => {
            successNotif.classList.replace('flex', 'hidden')
        }, timeout)
    } else {
        failedNotif.classList.replace('hidden', 'flex')
        setTimeout(() => {
            failedNotif.classList.replace('flex', 'hidden')
        }, timeout)
    }
}
