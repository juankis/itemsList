$(document).ready(function () {
    initialFunctions = () => {
        $("#itemsList").sortable();
        $("#itemsList").disableSelection();
    }

    updateItems = () =>  {
        $.ajax({
            url: "http://localhost:9000/get_items",
            type: 'GET',
            success: function (data) {
                console.log(data)
                fillItems(data.items)
            },
            cache: false,
            contentType: false,
            processData: false
        });
    }

    createItem = (picture, description, title, updated_at) => {
        return '<div class="card">'+
                    '<img class="card-img-top" src=http://localhost:9000/"'+picture+'" alt="Card image cap">'+
                    '<div class="card-body">'+
                    '<h4 class="card-title">'+title+'</h4>'+
                    '<p class="card-text">'+description+'</p>'+
                    '<p class="card-text">'+
                    '<small class="text-muted">'+updated_at+'</small>'+
                    '</p>'+
                    '</div>'+
                '</div>'
    }

    saveForm = () => {
        //e.preventDefault();   
        if(!$('form').valid())
            return false
        var form = $('form')[0];  
        var formData = new FormData(form);

        $.ajax({
            url: "http://localhost:9000/form_post",
            crossDomain: true,
            type: 'POST',
            data: formData,
            success: function (data) {
                console.log(data)
                $('form')[0].reset()
                $("#buttonCloseModal").click()
                updateItems()
            },
            cache: false,
            contentType: false,
            processData: false
        });
    }

    fillItems = (items) =>{
        $("#itemsList").empty();
        console.log("clear")
        items.forEach(item => {
            $("#itemsList").append(createItem(item.Picture, 
                                              item.Description,
                                              item.Title,
                                              "texto" 
                                            ))
        });
        $("#totalItems").html(items.length);
    }
    
    initialFunctions()
    updateItems()
});