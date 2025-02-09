document.onreadystatechange = function () {

	var ingredientId = 0;
	var ingredientCount = 0;

	function checkConnection() {
		var request = new XMLHttpRequest();
		request.open('GET', config.api+'/ping');
		request.onreadystatechange = function () {
			if (request.readyState == 4) {
				if (request.status != 200) {
					alert('There is no connection with the back end api');
				}
			}
		}

		request.setRequestHeader("Access-Control-Allow-Origin", "*")
		request.setRequestHeader("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		request.send();
	}

	function welcomeMessage() {
		var request = new XMLHttpRequest();
		request.open('GET', config.api+'/recipes/count');
		request.onreadystatechange = function () {
			if (request.readyState == 4) {
				if (request.status === 200) {
					var data = JSON.parse(request.responseText);

					if (data.message !== 'Not Found') {
						footer.innerHTML = `searching over ${data.count} recipes`
					}
				} else {
					footer.innerHTML = 'not found'
				}
			}
		}

		request.setRequestHeader("Access-Control-Allow-Origin", "*")
		request.setRequestHeader("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		request.send();
	}

	function showAll() {
		result.innerHTML = ''
		var request = new XMLHttpRequest();
		request.open('GET', config.api+'/recipes/');
		request.onreadystatechange = function () {
			if (request.readyState == 4) {
				if (request.status === 200) {
					var data = JSON.parse(request.responseText);

					if (data != null) {
						message.style.display = "none"
						result.style.display = "block"

						content = `<p>${data.length} results</p>`;
						content += '<table style="border-collapse: collapse;border-spacing: 0;">'
						content += '<thead>'
						content += '<tr>'
						content += '<th>Name</th>'
						content += '<th>Description</th>'
						content += '</tr>'
						content += '</thead>'
						content += '<tbody>'

						for (var i = 0; i < data.length; i++) {
							content += '<tr>'
							content += `<td><a href="vip.html?id=${data[i].id}">${data[i].name}</a></td>`
							content += `<td>${data[i].description}</td>`
							content += '</tr>'
						}

						content += '</tbody>'
						content += '</table>'

						result.innerHTML = content;
					} else {
						result.style.display = "none";
						message.style.display = "block";
						message.innerHTML = `no results for '${data.query}'`;
					}
				} else {
					result.style.display = "none";
					message.style.display = "block";
					message.innerHTML = 'there was an error';
				}
			}
		}

		request.setRequestHeader("Access-Control-Allow-Origin", "*")
		request.setRequestHeader("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		request.send();
	}

	function removeSearchIngredient(id) {
		var ingredient = document.getElementById('ingredient_' + id);
		ingredient.parentNode.removeChild(ingredient);
		ingredientCount--;
	}

	function addSearchIngredient() {
		if (ingredientCount < 3) {
			var addList = document.getElementById('ingredients');
			var docstyle = addList.style.display;
			if (docstyle == 'none') addList.style.display = '';

			ingredientId++;
			ingredientCount++;

			var text = document.createElement('div');
			text.id = 'ingredient_' + ingredientId;
			text.innerHTML = " \
        					<input id='name" + ingredientId + "' type='text' value='' placeholder='ingredient name' autocomplete='off' required/> \
        					<input id='remove"+ ingredientId + "' type='button' value='X' />";

			addList.appendChild(text);
			var i = ingredientId;
			document.getElementById("remove" + ingredientId).addEventListener("click", function () { removeSearchIngredient(i); }, false);
			document.getElementById("name" + ingredientId).focus();
		}
	}

	function search() {
		result.innerHTML = ''
		var query = document.querySelector('#search').value;
		var request = new XMLHttpRequest();

		var params = "";
		for (var i = 1; i <= ingredientId; i++) {
			if (document.getElementById("name" + i) != null) {
				if (i > 1 && params != "") {
					params += ",";
				}

				params += document.getElementById("name" + i).value
			}
		}

		request.open('GET', config.api+'/search?q=' + query + '&ingredient=' + params);
		request.onreadystatechange = function () {
			if (request.readyState == 4) {
				if (request.status === 200) {
					var data = JSON.parse(request.responseText);

					if (data.recipes != null) {
						message.style.display = "none"
						result.style.display = "block"

						if (params != '') {
							content = `<p>${data.recipes.length} results for '${data.query}', ingredients: '${params}'</p>`;
						} else {
							content = `<p>${data.recipes.length} results for '${data.query}'</p>`;
						}
						content += '<table style="border-collapse: collapse;border-spacing: 0;">'
						content += '<thead>'
						content += '<tr>'
						content += '<th>Name</th>'
						content += '<th>Description</th>'
						content += '</tr>'
						content += '</thead>'
						content += '<tbody>'

						for (var i = 0; i < data.recipes.length; i++) {
							content += '<tr>'
							content += `<td><a href="vip.html?id=${data.recipes[i].id}">${data.recipes[i].name}</a></td>`
							content += `<td>${data.recipes[i].description}</td>`
							content += '</tr>'
						}

						content += '</tbody>'
						content += '</table>'

						result.innerHTML = content;
					} else {
						result.style.display = "none";
						message.style.display = "block";
						message.innerHTML = `no results for '${data.query}'`;

						if (params != '') {
							message.innerHTML = `no results for '${data.query}', ingredients: '${params}`;
						} else {
							message.innerHTML = `no results for '${data.query}'`;
						}

					}
				} else {
					result.style.display = "none";
					message.style.display = "block";
					message.innerHTML = 'there was an error';
				}
			}
		}

		request.setRequestHeader("Access-Control-Allow-Origin", "*")
		request.setRequestHeader("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		request.send();
	}

	if (document.readyState === 'complete') {
		checkConnection();

		var result = document.querySelector('#result');
		var searchForm = document.querySelector('#search-form');
		document.getElementById("show-all").addEventListener("click", showAll, false);
		document.getElementById("add-ingredient").addEventListener("click", addSearchIngredient, false);

		welcomeMessage()

		searchForm.addEventListener('submit', function (e) {
			e.preventDefault()
			search()
		});
		
	    	showAll()

	}
}
