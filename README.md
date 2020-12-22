# District_Ward

Để sử dụng đầu tiên phải cài đặt gigonic, có thể cài đặt thông qua câu lệnh
go get "github.com/gin-gonic/gin"

Cấu trúc thư mục của District&Ward giống hệt nhau

Để có thể chạy, đầu tiên cấu hình lại mục const trong District_Ward/district/connect/postgres.go

Sau đó vào District_Ward/district/main.go
Trong terminal gõ: go run main.go

sever sẽ được chạy ở cổng 8080, có 2 phương thức là get và post

get:  127.0.0.1:8080/api/story/Read
post: 127.0.0.1:8080/api/story/Insert/0202

Sử dụng postman, phương thức post và đưa vào đường dẫn trên, chương trình sẽ tự tạo bảng district và thêm giá trị.

Sau đó sử dụng phương thức get với đường dẫn để xem kết quả.

Ở phần ward tương tự district
