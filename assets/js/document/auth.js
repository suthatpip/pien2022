var url = location.protocol + '//' + location.host;
  function confirmPasscode(){
    var obj= JSON.stringify({
        user: "aloha",
        email: "aloha@test.com" 
    })
 
    fetch(url + '/auth/passcode', {
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
            console.log(data.passcode);
            popup(data.passcode, "");
 
        })         
        .catch(err =>{
            console.log(err);
        }); 
}

async function popup(passcode, message){

 
    Swal.fire({
        title: 'กรอกรหัส 4 หลัก ที่ได้รับจากอีเมล์',
        inputLabel: 'Your age',
        
        html:        
        `<div class="flex-container">  
            <div class="input-group input-group-lg">              
                <input id="input1" type="text" class="form-control" minlength="1" maxlength="1" onkeypress='return event.charCode >= 48 && event.charCode <= 57'>
            </div>
        
            <div class="input-group input-group-lg">              
                <input id="input2" type="text" class="form-control" minlength="1" maxlength="1" onkeypress='return event.charCode >= 48 && event.charCode <= 57'>
            </div>
      
            <div class="input-group input-group-lg">              
                <input id="input3" type="text" class="form-control" minlength="1" maxlength="1" onkeypress='return event.charCode >= 48 && event.charCode <= 57'>
            </div>
       
            <div class="input-group input-group-lg">              
                <input id="input4" type="text" class="form-control" minlength="1" maxlength="1" onkeypress='return event.charCode >= 48 && event.charCode <= 57'>
            </div>
            <input type="hidden" id="passcode" value="` + passcode + `">
        </div>` ,
        
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
            if (message.length != 0) { 
                Swal.showValidationMessage(
                    message
                ) 
            }
        },
        preConfirm: () => {  
           let passcode=document.getElementById('passcode').value;        
           let d1=document.getElementById('input1').value;    
           let d2=document.getElementById('input2').value;    
           let d3=document.getElementById('input3').value; 
           let d4=document.getElementById('input4').value; 
           
           let code = d1.concat(d2, d3, d4);
           
            if (code.length != 4) { 
                Swal.showValidationMessage(
                    `กรอกรหัสไม่ครบ`
                ) 
            }else{ 
                return fetch(`/auth/passcode/`+ passcode +`/`+ code, {
                    method: 'POST',
                    headers: {
                    'Content-Type': 'application/json',
                    }
                })
                .then(response => {
                  if (!response.ok) {
                    throw new Error(response.statusText)
                  }
                  return response.json()
                })
                .catch(error => {
                  Swal.showValidationMessage(
                    `รหัสไม่ถูกต้อง`
                  )
                })
            } 
            return false;
        },
        // allowOutsideClick: () => !Swal.isLoading()
      }).then((result) => {  
          if (!result.isConfirmed) {
            console.log("OK");
          }else{
            console.log(result.passcode);
            popup(result.passcode, "รหัสไม่ถูกต้อง");
          }
      })

}

const getCode = (e) => {
    e = e || window.event;
    return e.key;
  };
  const handleKeyPress = (e) => {
    const key = getCode(e);
    if (isFinite(key)) {
      console.log(`Number ${key} was pressed!`);
    }
  };
  