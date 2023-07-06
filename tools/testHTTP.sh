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

# # Get Items
# curl -v -X GET "$host/items"
# # Get Item by ID
# curl -v -X GET "$host/items/id/24591"
# # Add Item by ID and Name
# curl -v -X POST "$host/items/id/3423/name/Lot_Z"
# # Delete a Item by ID
# curl -v -X DELETE "$host/items/id/3423"

# Get Bids
curl -X GET "$host/bids"
# Get Bid by UserID and ItemID
curl -X GET "$host/bids/54597/24595" # result amount should be 21
# Add Bid by UserID and ItemID and Amount
curl -X POST "$host/bids/54593/24595/100" # 54593 should now win the bid for 24595
# Check winner for 24595
curl -X GET "$host/winner/24595" # winner should be user 54593
# Update Bid of 54597 by UserID and ItemID and Amount to now win
curl -X PUT "$host/bids/54597/24595/105" # 54597 should now win the bid for 24595
# Check winner for 24595
curl -X GET "$host/winner/24595" # winner should be user 54597
# Delete Bid of 54597
curl -X DELETE "$host/bids/54597/24595" # winner should now be 54593 again
# Confirm by checking Bids on Item 24595
curl -X GET "$host/bids/24595" # winner should now be 54593 again
# Check all items 54593 has bid on
curl -X GET "$host/items/user/54593" # should only be on 24595

