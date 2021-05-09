function words(string) {
    return string.split(' ');
}

function sentence(ArrayOfString) {
    return ArrayOfString.join(' ');
}
function yell(string) {
    return string.toUpperCase();
}
function whisper(string) {
    return '*' + string.toLowerCase() + '*';
}
function capitalize(string) {
    return string.charAt(0).toUpperCase() + string.slice(1).toLowerCase();/*le premier char en uppercase, et le reste forme une autre chaîne de caractère en lwercase. Slice divise*/
}