var uploadDropzone = new Dropzone(document.body, {
    url: "/company/logo",
    thumbnailWidth: 80,
    thumbnailHeight: 80,
    parallelUploads: 1,
    maxFilesize: 1,
    maxFiles:1,
    acceptedFiles: "image/jpeg,image/png",
    previewTemplate: previewTemplate,
    autoQueue: false,
    previewsContainer: "#upload-previews",
    clickable: ".file-upload-input-button"
})

document.querySelector(".upload-start").onclick = function () { 
    uploadDropzone.enqueueFiles(uploadDropzone.getFilesWithStatus(Dropzone.ADDED))
 
}
  
document.querySelector(".upload-cancel").onclick = function () {
    uploadDropzone.removeAllFiles(true)
}
  
uploadDropzone.on("success", function (file, response) {
 
})

uploadDropzone.on("maxfilesexceeded", file => { 
    
});  

uploadDropzone.on("maxfilesreached", file => { 

}); 

uploadDropzone.on("error", file => {  
    uploadDropzone.removeAllFiles(true);

}); 

  

  // this.prototype.events = ["drop", "dragstart", "dragend", "dragenter", "dragover", 
  // "dragleave", "addedfile", "addedfiles", "removedfile", 
  // "thumbnail", "error", "errormultiple", "processing", 
  // "processingmultiple", "uploadprogress", "totaluploadprogress", 
  // "sending", "sendingmultiple", "success", "successmultiple", 
  // "canceled", "canceledmultiple", "complete", "completemultiple", 
  // "reset", "maxfilesexceeded", "maxfilesreached", "queuecomplete"];
    