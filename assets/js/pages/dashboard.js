var url = location.protocol + '//' + location.host;
$('#orderDetailModal').on('shown.bs.modal', function (e) {
   // var myBookId = $(this).data('id');
   var payment_code = $(e.relatedTarget).data('payment-code');
   var html=`<dl class="row text-sm" > `;

    fetch(url + '/dashboard/order/detail/'+payment_code, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        }
      })
      .then(response => response.json())
      .then(data => { 
        html +=`        
            <dt class="col-sm-3">ชื่อ-สกุล</dt>
            <dd class="col-sm-9">` + data.customer_name + `</dd> 

            <dt class="col-sm-3">บริษัท</dt>
            <dd class="col-sm-9">` + data.company.name + `</dd> 
            <dt class="col-sm-3">ที่อยู่</dt>
            <dd class="col-sm-9">` + data.company.address + `</dd> 
            <dt class="col-sm-3">เบอร์โทรศัพท์</dt>
            <dd class="col-sm-9">` + data.company.telephone + `</dd> 
            <dt class="col-sm-3">logo</dt>
            <dd class="col-sm-9">` + data.company.logo + `</dd> `

        html +=`<dt class="col-sm-12">  <hr class="border-3 opacity-75"></dt>`


        html +=`<dt class="col-sm-3">ออร์เดอร์</dt>
            <dd class="col-sm-9">` + data.order.order_no + `</dd>
            <dt class="col-sm-3">ใบกำกับภาษีหมายเลข</dt>
            <dd class="col-sm-9">` + data.order.tax_invoice_no + `</dd>  
            <dt class="col-sm-3">รหัสชำระเงิน</dt>
            <dd class="col-sm-9">` + data.order.payment_code + `</dd>  
            <dt class="col-sm-3">วันที่ชำระเงิน</dt>
            <dd class="col-sm-9">` + data.order.payment_date + `</dd> 
 
            <dt class="col-sm-3">วันที่เริ่ม</dt>
            <dd class="col-sm-9">` + data.publish_start_date + `</dd>  
            <dt class="col-sm-3">วันที่สิ้นสุด</dt>
            <dd class="col-sm-9">` + data.publish_end_date + `</dd>  
            `
        html +=`<dt class="col-sm-12">  <hr class="border-3 opacity-75"></dt>`
        data.products.forEach(products => {
            html +=`<dt class="col-sm-3">หมายเลข</dt>
            <dd class="col-sm-9">` + products.no + `</dd> 
            <dt class="col-sm-3">สินค้า</dt>
            <dd class="col-sm-9">` + products.product + `</dd> 
            <dt class="col-sm-3">ราคา </dt>
            <dd class="col-sm-9">` + products.product_baht + `</dd>`
            if(products.type=="file"){
                html +=`<dt class="col-sm-3">ไฟล์</dt>
                <dd class="col-sm-9">` + products.detail + `</dd> 
                <dt class="col-sm-3">ขนาด</dt>
                <dd class="col-sm-9">` + products.size + `</dd>` 
            }else{
                html +=`<dt class="col-sm-3">x</dt>
                <dd class="col-sm-9">` + products.detail + `</dd> ` 
               
            } 
        }) 
        html +=`</dl>`
        $('#order_detail').html(html);  
      }).catch((error) => {
        console.error('Error:', error)
      })

 
});


