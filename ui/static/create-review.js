let input = document.getElementById('user')

input.oninvalid = function(event) {
	event.target.setCustomValidity('Username should contain 6-30 latin letters. e.g. Anne Kim');
}
