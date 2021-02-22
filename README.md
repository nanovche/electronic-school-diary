# electronic-diary-web-app
Web application with Golang backend allowing users to manage their day to day schedule.

Core features and functionality of the web-app:

administrator ->

teacher -> 
  1. assess student
  2. write absence notes
  3. can be a class teacher
  
  // optioinal -> write behavioural(praise/critisize) notes
  
parent ->
  1. Associate himself with his child from (dropdown/radiobutton)menu of classes and then from another menu picks his child
  2. can see his child's(student) information (assessments, absence notes) 
  3. can upload sickleave note for his child as image/pdf
    
  Surfflow: logs in ->   
      
class ->

 /*optional associate with more than one child(has more kids)
 can revert the process and pick another child if he has made a mistake*/
 
 
 /* 1.some of way of communicating between the four user roles 
    2.real time bus tracking system(googlemaps api?/цгм?
    3.ajax ) */

Entities in database system:
Administrator ->
Teacher ->
Parent -> represents his child
Class -> comprises of many students 
Student
School -> comprises of many classes, many teachers, many students

