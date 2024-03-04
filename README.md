# simple-list-interview


API List คร่าวๆ
1. POST /login
2. GET /apis/interviews?lastItemOrder=0&&size=20
    ใช้ในการ fetch interviews list และใช้ตอน กด see more โดยการเปลี่ยน query 
3. GET /apis/interviews/:interviewId/detail
    ใช้ get detail ของการ click หน้า interview detail
4. PATCH /apis/interviews/:interviewId/detail
   ใช้ในการ update การเปลี่ยนแปลง เช่น edit title, description, status
5. GET /apis/interviews/:interviewId/history
   ใช้ในการ get history list ของ interview detail 
6. GET /apis/interviews/:interviewId/comments
    ใช้ในการ get comment list ใน interview นั้นค่ะ
7. PUT /interviews/:interviewId/comments/:commentId
   ใช้ในการ แก้ไข comment description



สิ่งที่สามารถ improve เพิ่มได้ใน project นี้คือ
1. unit test per service
2. แยก เส้น login ไปอีก domain นึงค่ะ
3. แยก middleware ไปเป็น lib กลาง
4. เพิ่มเส้น api สำหรับการ create, delete interview record
5. เพิ่มเส้น api สำหรับการ create, delete comment record