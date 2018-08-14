const searchBox = document.getElementById('input');
const resultContainer = document.querySelector('.results');
let timeout = null;

searchBox.onkeyup = function (event) {
	clearTimeout(timeout);
	timeout = setTimeout(function () {
		fetchUsers(searchBox.value);
	}, 600);
};

function fetchUsers(name) {
	fetch(`/users/${name}`)
		.then((resp) => resp.json())
		.then((users) => {
			renderUsers(users);
		})
		.catch((error) => {
			console.log(error);
		});
}

function renderUsers(users) {
	resultContainer.innerHTML = "";
	let htmlString = "";
	if (users.length > 0) {
		for (let index = 0; index < users.length; index++) {
			htmlString += `<li>${users[index]['Name']}</li>`;
		}
	}
	resultContainer.innerHTML = `<ul class="list">${htmlString}</ul>`;
}