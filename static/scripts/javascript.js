


window.onload=function(){

    function updateBoard(direct) {
        var $gridVals = [];
    
        $('.grid-item').each(function( index ) {
            //console.log( index + ": " + $( this ).text() );
            $gridVals.push($( this ).text());
          });
    
          var $respVals = JSON.stringify({direction: direct, values: $gridVals});
          console.log($respVals);
    
        console.log($gridVals);
        // send AJAX POST request
        var $urlString = "/"
        $.ajax({
            url: $urlString,
            type: 'POST',
            dataType : "json",
            contentType: "application/json",
            data: $respVals,
            //{'direction': 'right', 'values': $gridVals},
            success: function(response) {
                console.log("in success")
                console.log(JSON.stringify(response));
                $('.grid-item').each(function(index){
                    console.log(response[index])
                    $( this ).text(response[index]);
                });
    
                //console.log(response);
            },
            error: function(error){
                console.log(error);
                // $.ajax({
                //     url: $urlString,
                //     type: 'GET'
                //});
            }
        });
        }
        $('#slideRight').on('click', function(){
            //do someting
            console.log('right clicked')
            updateBoard("right");
        })
        $('#slideUp').on('click', function(){
            //do someting
            console.log('Up clicked')
            updateBoard("up");
        })
        $('#slideDown').on('click', function(){
            //do someting
            console.log('Up clicked')
            updateBoard("down");
        })
        $('#slideLeft').on('click', function(){
            //do someting
            console.log('Up clicked')
            updateBoard("left");
        })
    // try simple alert
    // $('#slideRight').on("click", function(){
    //     console.log("button clicked");
    //     //updateBoard("right");
    // )}


        

    // $('#slideUp').on("click", function(){
    //     (console.log("Up button clicked");
    //     //updateBoard("up");
    // )}

        // var $gridVals = [];


        // $('.grid-item').each(function( index ) {
        //     //console.log( index + ": " + $( this ).text() );
        //     $gridVals.push($( this ).text());
        //   });

        //   var $respVals = JSON.stringify({direction: 'up', values: $gridVals})
        //   console.log($respVals)

        // console.log($gridVals);
        // // send AJAX POST request
        // var $urlString = "/"
        // $.ajax({
        //     url: $urlString,
        //                 type: 'POST',
        //                 dataType : "json",
        //                 contentType: "application/json",
        //                 data: $respVals,
        //                 //{'direction': 'right', 'values': $gridVals},
        //                 success: function(response) {
        //                     console.log("in success")
        //                     console.log(JSON.stringify(response));
        //                     $('.grid-item').each(function(index){
        //                         console.log(response[index])
        //                         $( this ).text(response[index])
        //                     })

        //                     // repopulate board with response values


        //                     // if the response is a success, remove the table row
        //                     //$row.remove();
        //                     console.log(response);
        //                 },
        //                 error: function(error){
        //                     console.log(error);
        //                     // $.ajax({
        //                     //     url: $urlString,
        //                     //     type: 'GET'
        //                     //});
        //                 }
        //             });

    // })



        // $(".delBtn").on( "click", function() {
        //     // when user tries to delet a record, make them confirm
        //     if (confirm("Are you sure?")){
        //         console.log("confirmed delete");
        //         var $row = jQuery(this).closest('tr');
        //         //var $columns = $row.find('.serviceRecordID');
        //         var $recordToDelete = $row.find('.serviceRecordID');
        //         var $clientID = jQuery(this).closest('tr').find('.clientID').text();
        //         console.log("clientid: " + $clientID);
        //         var $value = $recordToDelete.text();
        //         console.log($value);
                
        //         // send ajax get with parameters to delete
        //         var $urlString = 'add_Service_' + $clientID;
        //         console.log("url: " + $urlString)
        //         $.ajax({
        //             url: $urlString,
        //             type: 'GET',
        //             data: {'recordID': $value,},
        //             success: function(response) {
        //                 // if the response is a success, remove the table row
        //                 $row.remove();
        //                 console.log(response);
        //             },
        //             error: function(error){
        //                 console.log(error);
        //                 $.ajax({
        //                     url: $urlString,
        //                     type: 'GET'
        //                 });
        //             }
        //         });
                            
        //     } else {
        //         console.log("user concelled");
        //     }
        // });
    }