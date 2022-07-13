var url = location.protocol + '//' + location.host;

document.addEventListener('DOMContentLoaded', function () {
    window.stepper = new Stepper(document.querySelector('.bs-stepper'))
})

function isEmpty(str) {
    return (!str || str.length === 0);
}

function hideProcessing(){
    console.log("hideProcessing");
    //  $("#processing").removeClass("in");
     // $(".modal-backdrop").remove();
      $("#processing").hide();
}
  
function processingError() {
    console.log("processingError");
    return ` 
    <div class="card">
      <div class="card-body"> 
        <br>
        <p class="card-text text-danger"> พบข้อผิดพลาด กรุณาลองใหม่อีกครั้ง :(</p>             
      </div>
    </div> 
    `;
}
