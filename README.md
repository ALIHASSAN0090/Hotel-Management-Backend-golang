# **Hotel Management Backend**  
### **Optimized Backend Solution for Hotel Management**

This project is a robust, feature-rich backend application built with **Golang**, optimized for high performance and low latency. It integrates advanced tools and techniques to streamline hotel management operations, provide real-time insights, and enhance the overall user experience.

---

## **Features**
- **Admin Panel:**  
  - Real-time sales analysis and reporting.  
  - Role-based access control for admins and users.  
  - AI-driven recommendations based on past booking and sales data.  

- **AI-Powered Analytics:**  
  - Advanced monthly sales analysis powered by AI, enabling smarter business decisions.  

- **Customer Support:**  
  - Integrated AI chat assistant for 24/7 customer support to handle common queries and enhance customer engagement.  

- **Task Automation:**  
  - Cron jobs for routine task automation (e.g., scheduled report generation, periodic database cleanup).  

- **Booking and Room Management:**  
  - APIs to manage room availability, bookings, and cancellations.  

- **Email Notifications:**  
  - SMTP-based email functionality for order confirmations and notifications.  

- **Performance Optimization:**  
  - Leveraged Go routines for concurrent processing.  
  - Redis caching for frequently accessed resources.  

- **Database Management:**  
  - Complex PostgreSQL schema for detailed hotel data.  
  - Custom database migrator for smooth and efficient updates.  

---

## **Tech Stack**
- **Backend:** Golang  
- **Database:** PostgreSQL  
- **Caching:** Redis  
- **AI Integration:** OpenAI API for AI-driven analytics and chat assistant  
- **Deployment:** Docker, Kubernetes  
- **CI/CD:** GitHub Actions  

---

## **Schema Overview**
![Sample Image](schema-img/schema.png)  
*(This image represents the PostgreSQL schema structure used for the project.)*

---

## **Installation and Usage**

### **1. Clone the Repository**
```bash
git clone https://github.com/alihassan0090/hotel-management-backend.git
cd hotel-management-backend
```

### **2. Set Up Environment Variables**
Create a `.env` file and provide the necessary configurations:
```env
PORT=your_server_port
DB_HOST=your_database_host
DB_PORT=your_database_port
DB_USER=your_database_user
DB_NAME=your_database_name
DB_PASSWORD=your_database_password
ADMIN_DEFAULT_EMAIL=default_admin_email
ADMIN_DEFAULT_PASSWORD=default_admin_password
ADMIN_ROLE=admin_role_name
JWT_SECRET=your_jwt_secret
MAIL_HOST=smtp_mail_host
MAIL_PORT=smtp_mail_port
MAIL_USERNAME=smtp_mail_username
MAIL_PASSWORD=smtp_mail_password
OPENAI_KEY=your_openai_api_key
AWS_SECRET_ACCESS_KEY=your_aws_secret_key
AWS_ACCESS_KEY=your_aws_access_key
REGION=aws_region
ECR_REPOSITORY=aws_ecr_repository
ECS_CLUSTER=aws_ecs_cluster_name
ECS_SERVICE=aws_ecs_service_name
ECS_TASK_DEFINITION=aws_task_definition_name
```

### **3. Run the Application**
#### Using Docker:
```bash
make docker-setup
```

#### Locally:
```bash
make local-setup
```

---

## **Deployment**

### **CI/CD Pipeline**
- Integrated **GitHub Actions** for automated testing, build, and deployment workflows.

### **Kubernetes Deployment**
- The application is containerized with **Docker** and deployed on a **Kubernetes** cluster for scalability and reliability.

---

## **Project Highlights**
- **Duration:** 3 months (completed alongside full-time employment).  
- Combines AI technologies and Golang to solve real-world hotel management challenges effectively.  
- Implements modern backend development practices to ensure scalability, maintainability, and performance.  

---

## **Contributing**
We welcome contributions to make this project even better! Feel free to submit issues or pull requests.  

---

## **License**
This project is licensed under the [MIT License](LICENSE).  

---

## **Contact**
For any inquiries or support:  
📧 Email: [alihassankhan285@gmail.com](mailto:alihassankhan285@gmail.com)  
🌐 GitHub: [alihassan0090](https://github.com/alihassan0090)  

---
