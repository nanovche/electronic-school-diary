# electronic-diary-web-app
Web application with Golang backend, golang's builtin html templates, gorm(go's orm) to communicate with database

Core features and functionality of the web-app:

all -> change data accoring to user roles
all -> register/login/logout

administrator -> 
  1. CRUD student
  2. CRUD teacher
  3. CRUD subjects
  4. CRUD classes
  5. link the system (create classes with students and add subjects to those a classes; link teachers to subjects/classes) 

teacher -> 
  1. CRUD assess student/students
  2. CREATE/UPDATE/DELETE/READ -> write absence notes
  3. CREATE/UPDATE/DELETE/READ -> write behavioural(praise/critisize) notes
  4. can be a class teacher
  5. communicate with parent
  6. READ his students' info
  
parent ->
  1. Associate himself with his child from (dropdown/radiobutton)menu of classes and then from another menu picks his child
  2. READ -> can see his child's(student) information (assessments, absence notes) 
  3. CREATE? -> can upload sickleave note for his child as image/pdf
  4. CREATE? -> communicate with teacher
    
 /* 1.optional associate with more than one child(has more kids)
    can revert the process and pick another child if he has made a mistake
  
    2.some of way of communicating between the four user roles 
    3.real time bus tracking system(googlemaps api?/цгм?
 */

Entities in database system:
Administrator
Teacher
Subject
Parent
Class 
Student
sickleave note
absence note
behavioural note
mark note

