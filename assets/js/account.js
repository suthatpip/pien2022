  
  var url = location.protocol + '//' + location.host;
 var companyDropzone;
 var previewNode;
 var previewTemplate;

 function edit_company(name, telephone, tax_address, logo ,about, code){
  
  Swal.fire({
    title: 'เปลี่ยนแปลงข้อมูล ' ,    
    html:        
    `<img src="`+ logo +` " alt="" data-dz-thumbnail />
    <hr>
      <div class="flex-container">  
          <div class="card-body">            
            <div class="input-group mb-3">
            <div class="table table-striped files">
                <div class="row">
                  <div class="col-7 bg-gray-light">
                    <div class="table table-striped files" id="previews">
                      <div id="_logo" class="row mt-0"> 
                        <div class="col-auto ">
                          <span class="preview"> 
                            <img src="data:," alt="" data-dz-thumbnail />
                          </span>
                        </div>
                      </div>
                    </div> 
                  </div> 
                  <div id="actions" class="col-5">
                    <div class="btn-group  d-flex align-items-start">
                      <span class="btn btn-success col fileinput-button">
                        <i class="fas fa-plus"></i>
                        <span class="text-sm">เปลี่ยนโลโก้</span>
                      </span>
                      <button data-dz-remove class="btn btn-danger clear">
                        <i class="fas fa-trash"></i>
                      </button>  
                    </div>
                  </div>
                </div>                 
              </div>
            </div> 
              <div class="form-group">
                <label class="float-left">ชื่อ นามสกุล</label>
                <input id="_name" type="email" class="form-control form-control-sm" maxlength="150" placeholder="`+ name +`">
              </div>              
              <div class="form-group">
                <label class="float-left">หมายเลขโทรศัพท์</label>          
                <div class="input-group">
                <div class="input-group-prepend">
                  <span class="input-group-text"><i class="fas fa-phone"></i></span>
                </div>
                  <input id="_telephone" type="text" class="form-control form-control-sm"  maxlength="30" inputmode="text" placeholder="`+ telephone +`">
                </div>
              </div>
              <div class="form-group">
                <label class="float-left">ที่อยู่สำหรับออกใบกำกับภาษี</label>
                <textarea id="_tax_address" class="form-control form-control-sm" rows="3"  maxlength="300" placeholder="`+ tax_address +`"></textarea>
              </div>
              <div class="form-group">
                <label class="float-left">หมายเหตุ</label>
                <textarea id="_about" class="form-control form-control-sm" rows="3" maxlength="100" placeholder="`+ about +`"></textarea>
              </div>                    
          </div>    
      </div>`  ,
    
    inputAttributes: {
      autocapitalize: 'off'
    },
    showCancelButton: true,
    cancelButtonText:"ยกเลิก",
    confirmButtonText: 'ยืนยัน',
    confirmButtonColor:"#007bff",
    allowOutsideClick: false,
    showLoaderOnConfirm: true,
   
    didOpen: () => { 
      init(code); 
    },  
    
    preConfirm: () => {  
        
        companyDropzone.enqueueFiles(companyDropzone.getFilesWithStatus(Dropzone.ADDED))
        
        if (document.getElementById('_name').value.length > 0) {
          name= document.getElementById('_name').value;
         }

         if (document.getElementById('_telephone').value.length > 0) {
          telephone= document.getElementById('_telephone').value;
         }

         if (document.getElementById('_tax_address').value.length > 0) {
          tax_address= document.getElementById('_tax_address').value;
         }

         if (document.getElementById('_about').value.length > 0) {
          about= document.getElementById('_about').value;
         } 

        var obj= JSON.stringify({ 
          name: name,
          address: tax_address,
          telephone: telephone,
          about: about, 
          code: code,
      })
   
      fetch(url + '/account/company' , {
              method: 'POST',
              headers: {
              'Content-Type': 'application/json',
              },
              body: obj
          })
          .then(function (response) {
              if (!response.ok) {  
                  throw new Error("HTTP status " + response.status);
              }   
              return response.json();
          })
          .then(function (data) { 
            refresh(); 
          })         
          .catch(err =>{
              console.log(err);
          });  
    }, 
  }).then((result) => { 
    console.log("result");
  })
 }

function init(code){
  Dropzone.autoDiscover = false;
  previewNode = document.querySelector("#_logo");
  previewNode.id = code;
  previewTemplate =  previewNode.parentNode.innerHTML;
  previewNode.parentNode.removeChild(previewNode);
    companyDropzone = new Dropzone(document.body, {
    url: "/account/" + code +"/logo",
    thumbnailWidth: 80,
    thumbnailHeight: 80,
    parallelUploads: 1,
    maxFilesize: 1,
    maxFiles: 1,
    acceptedFiles: "image/jpeg,image/png",
    previewTemplate:  previewTemplate,
    autoQueue: false,
    previewsContainer: "#previews",
    clickable: ".fileinput-button"
  })
 
 
  document.querySelector(".clear").onclick = function () {  
      companyDropzone.removeAllFiles(true) 
  } 
    // this.prototype.events = ["drop", "dragstart", "dragend", "dragenter", "dragover", 
    // "dragleave", "addedfile", "addedfiles", "removedfile", 
    // "thumbnail", "error", "errormultiple", "processing", 
    // "processingmultiple", "uploadprogress", "totaluploadprogress", 
    // "sending", "sendingmultiple", "success", "successmultiple", 
    // "canceled", "canceledmultiple", "complete", "completemultiple", 
    // "reset", "maxfilesexceeded", "maxfilesreached", "queuecomplete"];
    companyDropzone.on("addedfile", file => { 
      console.log("addedfile");  
    });  
    companyDropzone.on("complete", file => { 
      console.log("complete");  
    });
    companyDropzone.on("success", file => { 
      console.log("success");  
    });
    companyDropzone.on("maxfilesexceeded", file => { 
      companyDropzone.removeAllFiles(true)
      companyDropzone.addFile(file); 
    });   
    
    companyDropzone.on("error", file => {  
      companyDropzone.removeAllFiles(true); 
    }); 
}
 
 function delete_company(code, company){
 
    Swal.fire({
        title: 'คุณแน่ใจใช่ไหม?',
        text: "กด ยืนยัน เพื่อลบข้อมูล",
        icon: 'warning',
        showCancelButton: true,

        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        cancelButtonText:"ยกเลิก",
        confirmButtonText: 'ยืนยันการลบ '
      }).then((result) => {
        if (result.isConfirmed) {         
              fetch(url + '/account/company/del/'+ code , {
                  method: 'POST',
                  headers: {
                  'Content-Type': 'application/json',
                  }  
              })
              .then(function (response) {
                  if (!response.ok) {  
                      throw new Error("HTTP status " + response.status);
                  }   
                  return response.json();
              })
              .then(function (data) { 
                
               if (data.status=="OK"){ 
                  refresh();
               }else{
                Swal.fire({
                  icon: 'error',
                  title: 'Oops...',
                  text: 'พบความผิดพลาดบางอย่าง',
                  
                })
               }
                
              })         
              .catch(err =>{
                  console.log(err);
              }); 

         
        }
        
      })
 }
 
 function new_company(){

  Swal.fire({
    title: 'เริ่มกันเลย' ,    
    html:        
    `<hr>
      <div class="flex-container">  
          <div class="card-body">            
            <div class="input-group mb-3">
            <div class="table table-striped files">
                <div class="row">
                  <div class="col-7 bg-gray-light">
                    <div class="table table-striped files" id="previews">
                      <div id="_logo" class="row mt-0"> 
                        <div class="col-auto ">
                          <span class="preview"> 
                            <img src="data:," alt="" data-dz-thumbnail />
                          </span>
                        </div>
                      </div>
                    </div> 
                  </div> 
                  <div id="actions" class="col-5">
                    <div class="btn-group  d-flex align-items-start">
                      <span class="btn btn-success col fileinput-button">
                        <i class="fas fa-plus"></i>
                        <span class="text-sm">โลโก้</span>
                      </span>
                      <button data-dz-remove class="btn btn-danger clear">
                        <i class="fas fa-trash"></i>
                      </button>  
                    </div>
                  </div>
                </div>                 
              </div>
            </div> 
              <div class="form-group">
                <label class="float-left">ชื่อ นามสกุล<code>*</code></label>
                <input id="_name" type="email" class="form-control form-control-sm" maxlength="150" placeholder=""  >
              </div>              
              <div class="form-group">
                <label class="float-left">หมายเลขโทรศัพท์<code>*</code></label>          
                <div class="input-group">
                <div class="input-group-prepend">
                  <span class="input-group-text"><i class="fas fa-phone"></i></span>
                </div>
                  <input id="_telephone" type="text" class="form-control form-control-sm"  maxlength="30" inputmode="text" placeholder="">
                </div>
              </div>
              <div class="form-group">
                <label class="float-left">ที่อยู่สำหรับออกใบกำกับภาษี<code>*</code></label>
                <textarea id="_tax_address" class="form-control form-control-sm" rows="3"  maxlength="300" placeholder=""></textarea>
              </div>
              <div class="form-group">
                <label class="float-left">หมายเหตุ</label>
                <textarea id="_about" class="form-control form-control-sm" rows="3" maxlength="100" placeholder=""></textarea>
              </div>                    
          </div> 
          <input type="hidden" id="_code" value=""> 
      </div>`  ,
    
    inputAttributes: {
      autocapitalize: 'off'
    },
    showCancelButton: true,
    cancelButtonText:"ยกเลิก",
    confirmButtonText: 'ยืนยัน',
    confirmButtonColor:"#007bff",
    allowOutsideClick: false,
    showLoaderOnConfirm: true,
    inputValidator: (value) => {
      if (!value) {
        return 'You need to write something!'
      }
      
    },
    didOpen: () => { 
       initcode();         
    },  
    didClose : () => {         
      companyDropzone=null;
   },  
   
    preConfirm: () => {   
      var name= document.getElementById('_name').value;          
      var telephone= document.getElementById('_telephone').value;         
      var tax_address= document.getElementById('_tax_address').value;  
      var about= document.getElementById('_about').value;
      var code= document.getElementById('_code').value;

      if (name.length ==0 || telephone.length ==0 ||tax_address.length ==0  ){
        Swal.showValidationMessage('ข้อมูลไม่ครบถ้วน กรุณาเช็คข้อมูลที่มีเครื่องหมาย <code>*</code> ปรากฎอยู่') 
        return;
      } 


      var obj= JSON.stringify({ 
          name: name,
          address: tax_address,
          telephone: telephone,
          about: about, 
          code: code,
      })
   
      fetch(url + '/account/company/new' , {
          method: 'POST',
          headers: {
          'Content-Type': 'application/json',
          },
          body: obj
      })
      .then(function (response) {
          if (!response.ok) {  
              throw new Error("HTTP status " + response.status);
          }   
          return response.json();
      })
      .then(function (data) {
        companyDropzone.enqueueFiles(companyDropzone.getFilesWithStatus(Dropzone.ADDED))  
        refresh();
      })         
      .catch(err =>{
          console.log(err);
      }); 
            
    },
    
  }) 
 }

 function initcode() {  
  fetch(url + '/account/company/init' , {
      method: 'POST',
      headers: {
      'Content-Type': 'application/json',
      } 
  })
  .then(function (response) {
      if (!response.ok) {  
          throw new Error("HTTP status " + response.status);
      }   
      return response.json();
  })
  .then(function (data) { 
    document.getElementById('_code').value=data.code;
    init(data.code);
    
  })         
  .catch(err =>{
      console.log(err);
  }); 
 }

 function refresh(){
  window.location.href ="/account"
 }
 