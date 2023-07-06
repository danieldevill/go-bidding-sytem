# Shell commands to test each HTTP method
# Based on: https://owasp.org/www-project-web-security-testing-guide/latest/4-Web_Application_Security_Testing/02-Configuration_and_Deployment_Management_Testing/06-Test_HTTP_Methods

host=http://localhost:3000

# # Get Users
# curl -v -X GET "$host/users"

# # Get User by ID
# curl -v -X GET "$host/users/id/54591"

# # Add User by ID and Name
# curl -v -X POST "$host/users/id/23/name/DAN"

# # Delete a User by ID
# curl -v -X DELETE "$host/users/id/23"

# Get Items
curl -v -X GET "$host/items"

# Get Item by ID
curl -v -X GET "$host/items/id/24591"

# Add Item by ID and Name
curl -v -X POST "$host/items/id/3423/name/Lot_Z"

# Delete a Item by ID
curl -v -X DELETE "$host/items/id/3423"