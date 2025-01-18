document.onreadystatechange = function () {

	function getParameterByName(name, url) {
		if (!url) url = window.location.href;
		name = name.replace(/[\[\]]/g, '\\$&');
		var regex = new RegExp('[?&]' + name + '(=([^&#]*)|&|#|$)'),
			results = regex.exec(url);
		if (!results) return null;
		if (!results[2]) return '';
		return decodeURIComponent(results[2].replace(/\+/g, ' '));
	}

	function showRecipe() {
		var id = getParameterByName('id');

		var request = new XMLHttpRequest();
		request.open('GET', config.api+'/recipes/' + id);
		request.onreadystatechange = function () {
			if (request.readyState == 4) {
				if (request.status === 200) {
					var data = JSON.parse(request.responseText);

					name.innerHTML = `${data.name}`;
					description.innerHTML = `${data.description}`;

					var ing = "";
					for (var j = 0; j < data.ingredients.length; j++) {
						i = data.ingredients[j];
						if(i.note == "") {
							if(i.unit == "a gusto") {
								ing += "<li> <strong>" + `${i.name}` + "</strong>, " + `${i.unit}` + "</li>";
							} else {
								ing += "<li> <strong>" + `${i.name}` + "</strong>, " + `${i.quantity}` + " " + `${i.unit}` + "</li>";
							}
						} else {
							ing += "<li> <strong>" + `${i.name}` + "</strong>, " + `${i.quantity}` + " " + `${i.unit}` + " (" + `${i.note}` + ") </li>"
						}
					}

					ingredients.innerHTML = ing
					directions.innerHTML = `${data.direction}`;

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
		var name = document.getElementById("name")
		var description = document.getElementById("description")
		var ingredients = document.getElementById("ingredients")
		var directions = document.getElementById("directions")

		showRecipe();
	}
}