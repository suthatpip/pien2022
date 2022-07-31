var url = location.protocol + '//' + location.host;
  function confirmPasscode(){
    var obj= JSON.stringify({
        user: "aloha",
        email: "aloha@test.com" 
    })
 
    fetch(url + '/auth/newpasscode', {
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
        ` <p class="text-xs text-secondary text-center">code: ` + passcode + `</p>  
        <div class="flex-container">  
            <div class="input-group input-group-lg">              
                <input id="input1"  type="text" focus class="form-control _passcode" minlength="1" maxlength="1"  >
                       
                <input id="input2" type="text" class="form-control _passcode" minlength="1" maxlength="1"   >
                       
                <input id="input3" type="text" class="form-control _passcode" minlength="1" maxlength="1" >
                   
                <input id="input4" type="text" class="form-control _passcode" minlength="1" maxlength="1" >
            </div>
            <input type="hidden" id="passcode" value="` + passcode + `">
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
        willOpen:() =>{ 

          const inputElements = document.getElementsByClassName('_passcode');
          for (let inputElement of inputElements) {           
              inputElement.addEventListener('keydown',enforceFormat);
              inputElement.addEventListener('keyup',nextElement);
          } 
        },
        didOpen: () => {
            if (message.length != 0) { 
                Swal.showValidationMessage(
                    message
                );
               
            }
            document.getElementById("input1").focus();
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
                clear();                
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
                    `พบความผิดพลาด`
                  )                 
                  clear();
                })
            } 
            return false;
        },
        // allowOutsideClick: () => !Swal.isLoading()
      }).then((result) => { 
        if (!result.isDismissed)  { 
          if (result.isConfirmed) {            
            if(result.value.result=="INVALID"){               
              popup(result.value.passcode, "รหัสไม่ถูกต้อง"); 
            } else if(result.value.result=="BLOCK"){  
              Swal.fire({
                icon: 'error',
                title: 'ระบบยกเลิกรหัสนี้',
                text: '!!คุณสามารถขอรหัสใหม่ได้หลังจากนี้!!', 
              })
            } else if(result.value.result=="VALID"){              
              auth(result.value.passcode, result.value.confirm);
            }else{
              popup(result.value.passcode, "พบความผิดพลาด"); 
            }      
          } 
        }          
      })

} 
function auth(passcode, confirm){
  console.log(passcode,'  === ', confirm);  
  window.location.href = "/auth/ready/" + passcode + '/'+ confirm;
}

function clear(){ 
  const inputElements = document.getElementsByClassName('_passcode');
  for (let inputElement of inputElements) {
     
      inputElement.value="";
  }
  document.getElementById("input1").focus();
}

const isNumericInput = (event) => {
	const key = event.keyCode; 
	return ((key >= 48 && key <= 57)|| (key >= 96 && key <= 105));
   
};
 
const enforceFormat = (event) => { 
    if(!isNumericInput(event)){       
        event.preventDefault();
    }
};

const nextElement = (event) => {
  var  ob = event.target;
  if(isNumericInput(event)){ 
    if(ob.nextElementSibling != null){   
      ob.nextElementSibling.focus();
    }else {            
      
      ob.parentNode.children[0].focus();
    }
  }
 
}


