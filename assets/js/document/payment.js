 
function initpayment(obj) { 

    fetch(url + '/payment/init', {
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
        
        $('#_company_name').text(data.company.name);
  
        $('#_name').text(data.customer_name);
        $('#_address').text(data.company.address);
        $('#_telephone').text(data.company.telephone);
        $('#_customer_code').text(data.company.code);
        $('#_company_logo').attr("src", data.company.logo);
  
        $('#_order').text(data.order.order_no);
        $('#_payment_due_date').text(data.order.payment_due);
        $('#_tax_invoice_no').text(data.order.tax_invoice_no);
        $('#_payment_code').text(data.order.payment_code);
        var html=``
        data.product.forEach(obj => {
 
          html +=`
          <tr class="text-sm">
            <td><span>` + obj.no + `</span></td>
            <td><span>` + obj.product + `</span></td>
            <td><span>` + obj.start_date + `</span></td>
            <td><span>` + obj.end_date + `</span></td>
            <td><span>` + obj.days + `</span></td>
            <td><span>` + obj.product_baht + `</span></td>
          </tr> `

        })
        $('#summary_products').html(html);      
        $('#_product_baht').text(data.product.product_baht);
        $('#_sub_total').text(data.sub_total);
        $('#_vat').text(data.vat);
        $('#_total').text(data.total);  
        
        stepper.next();
        
        // $("#processing").modal('hide'); 
      })
      .catch(err =>{
          //  $('#processing').modal('hide'); 
      });

}


function submitOrder(){
 
    var payment_code=$('#_payment_code').text();
     
    window.open(url+"/paysolution/" + payment_code, "_self"); 
   
}
 