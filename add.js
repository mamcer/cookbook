document.onreadystatechange = function () {

    var ingredientId = 0;

    function removeIngredient(id) {
        var ingredient = document.getElementById('ingredient_' + id);
        ingredient.parentNode.removeChild(ingredient);
    }

    function addIngredient() {
        var addList = document.getElementById('ingredients');
        var docstyle = addList.style.display;
        if (docstyle == 'none') addList.style.display = '';

        ingredientId++;

        var text = document.createElement('div');
        text.id = 'ingredient_' + ingredientId;
        text.innerHTML = " \
        <input id='name" + ingredientId + "' type='text' value='' placeholder='name' autocomplete='off' required/> \
        <input id='quantity" + ingredientId + "' type='number' step='0.1' value='1' placeholder='quantity' autocomplete='off' required /> \
        <input id='unit"+ ingredientId + "' type='text' value='' placeholder='unit' autocomplete='off' required /> \
        <input id='note"+ ingredientId + "' type='text' value='' placeholder='note' autocomplete='off' /> \
        <input id='remove"+ ingredientId + "' type='button' value='X' />";

        addList.appendChild(text);
        var i = ingredientId;
        document.getElementById("remove" + ingredientId).addEventListener("click", function () { removeIngredient(i); }, false);
        document.getElementById("name" + ingredientId).focus();
    }

    function cancel() {
        if (confirm('Are you sure? any change will be lost')) {
            location.href = 'index.html'
        }
    }

    function add() {
        var request = new XMLHttpRequest();
        request.open('POST', config.api+'/recipes', true);
        var rn = document.querySelector('#name').value;
        var rde = document.querySelector('#description').value;
        var lines = document.querySelector('#direction').value.split('\n');
        var rdi = lines.join("<br>")

        var params = `{"name":"${rn}","description":"${rde}","direction":"${rdi}","ingredients":[`;
        for (var i = 1; i <= ingredientId; i++) {
            if (document.getElementById("name" + i) != null) {
                if (i > 1) {
                    params += ",";
                }

                params += `{"name":"${document.getElementById("name" + i).value}",`
                params += `"quantity": ${parseFloat(document.getElementById("quantity" + i).value)},`
                params += `"unit":"${document.getElementById("unit" + i).value}",`
                params += `"note":"${document.getElementById("note" + i).value}"}`
            }
        }
        params += "]}";
        request.onreadystatechange = function () {
            if (request.readyState == 4) {
                if (request.status === 201) {
                    message.innerHTML = 'recipe successfully created!';
                    backToRecipes.style.display = "inline"
                    another.style.display = "inline"
                } else {
                    message.innerHTML = 'there was an error';
                }
            }
        }

        request.setRequestHeader("Access-Control-Allow-Origin", "*")
        request.setRequestHeader("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
        addForm.style.display = "none"
        message.style.display = "block"
        message.innerHTML = 'adding recipe...';
        request.send(params);
    }

    if (document.readyState === 'complete') {
        var addForm = document.querySelector('#add-form');
        var backToRecipes = document.querySelector('#back-to-recipes');

        document.getElementById("add-ingredient").addEventListener("click", addIngredient, false);
        document.getElementById("cancel").addEventListener("click", cancel, false);

        addIngredient();

        document.getElementById("name").focus();
        
        addForm.addEventListener('submit', function (e) {
            e.preventDefault()
            add()
        });
    }
}