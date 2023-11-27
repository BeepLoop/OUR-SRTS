const items = document.getElementById('sidebar-ul')

paths = window.location.href.split('/')
path = paths.pop()
if (path === '') {
    path = paths.pop()
}
console.log({ path })

for (const e of items.children) {
    if (e.id === path) {
        e.classList.add('highlight')
    } else {
        e.classList.remove('highlight')
    }
}
