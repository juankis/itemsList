$(document).ready(function () {

    var items

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
                items = data.items
                fillItems(items)
            },
            cache: false,
            contentType: false,
            processData: false
        });
    }

    createItem = (id,picture, description, title, updated_at) => {
        return '<div class="card">'+
                    '<img class="card-img-top" src=http://localhost:9000/pictures/'+picture+' alt="Card image cap">'+
                    '<div class="card-body">'+
                    '<h4 class="card-title">'+title+'</h4>'+
                    '<p class="card-text">'+description+'</p>'+
                    '<p class="card-text">'+
                    '<small class="text-muted">'+updated_at+'</small>'+
                    '<button type="button" onclick="fillItem('+id+')" class="btn btn-link btn-sm">Edit</button>'+
                    '<button type="button" onclick="deleteItem('+id+')" class="btn btn-link btn-sm">Delete</button>'+
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
        items.forEach(item => {
            $("#itemsList").append(createItem(item.Id, 
                                              item.Picture, 
                                              item.Description,
                                              item.Title,
                                              "texto" 
                                            ))
        });
        $("#totalItems").html(items.length);
    }
    
    deleteItem = (id) =>  {
        $.ajax({
            url: "http://localhost:9000/delete_item?id="+id,
            type: 'GET',
            success: function (data) {
                console.log(data)
                updateItems()
            },
            cache: false,
            contentType: false,
            processData: false
        });
    }

    edit = (id) =>  {
        $.ajax({
            url: "http://localhost:9000/edit_item",
            type: 'POST',
            success: function (data) {
                console.log(data)
                fillItems(data.items)
            },
            cache: false,
            contentType: false,
            processData: false
        });
    }

    fillItem = (id) => {
        $('#buttonModal').click()
        item = getItem(id)
        console.log(item)
        $("#title").val(item.Title)
        $("#description").val(item.Description)
        //$("#picture").val(item.Picture)   
    }
  
    getItem = (id) => {
        for (var i = 0; i < items.length; i++){
            if (items[i].Id == id){
               return items[i]
            }
        }
        return null
    }

    initialFunctions()
    updateItems()
});