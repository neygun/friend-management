# Users
* User object
```
{
  id: int64
  email: string
  created_at: time.Time
  updated_at: time.Time
}
```
**POST /friend-connection**
----
  Create a friend connection between two email addresses.
* **URL Params**  
  None
* **Headers**  
  Content-Type: application/json  
* **Data Params**  
```
  {
    "friends": [
        <user_email>,
        <user_email>
    ]
  }
```
* **Success Response:**  
* **Code:** 200  
  **Content:**
  ```
    {"success": true}
  ```
* **Error Response:**  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Invalid email."
    }
    ```
  OR  
  * **Code:** 404  
  **Content:**
    ```
    {
        "code": 404,
        "description": "Email not found."
    }
    ```
  OR  
  * **Code:** 409  
  **Content:**
    ```
    {
        "code": 409,
        "description": "Friend connection already exists."
    }
    ```
  OR  
  * **Code:** 409  
  **Content:**
    ```
    {
        "code": 409,
        "description": "Blocking relationship exists."
    }
    ```
  OR  
  * **Code:** 500  
  **Content:**
    ```
    {
        "code": 500,
        "description": "Internal server error."
    }
    ```

**GET /friends-list/:email**
----
  Retrieve the friends list for an email address.
* **URL Params**  
  *Required:* `email=[string]`
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json
* **Success Response:** 
* **Code:** 200  
  **Content:**
  ```
  {
    "success": true,
    "friends": [
        <user_email>,
        <user_email>,
        <user_email>
    ],
    "count": integer
  }
  ```
* **Error Response:**  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Invalid email."
    }
    ```
  OR  
  * **Code:** 404  
  **Content:**
    ```
    {
        "code": 404,
        "description": "Email not found."
    }
    ```
  OR  
  * **Code:** 500  
  **Content:**
    ```
    {
        "code": 500,
        "description": "Internal server error."
    }
    ```

**POST /common-friends**
----
  Retrieve the common friends list between two email addresses.
* **URL Params**  
  None
* **Headers**  
  Content-Type: application/json  
* **Data Params**  
```
  {
    "friends": [
        <user_email>,
        <user_email>
    ]
  }
```
* **Success Response:**  
* **Code:** 200  
  **Content:**
  ```
  {
    "success": true,
    "friends": [
        <user_email>,
        <user_email>,
        <user_email>
    ],
    "count": integer
  }
  ```
* **Error Response:**  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Invalid email."
    }
    ```
  OR  
  * **Code:** 404  
  **Content:**
    ```
    {
        "code": 404,
        "description": "Email not found."
    }
    ```
  OR  
  * **Code:** 500  
  **Content:**
    ```
    {
        "code": 500,
        "description": "Internal server error."
    }
    ```

**POST /subscribe**
----
  Subscribe to updates from an email address.
* **URL Params**  
  None
* **Headers**  
  Content-Type: application/json  
* **Data Params**  
```
  {
    "requestor": <user_email>,
    "target": <user_email>
  }
```
* **Success Response:**  
* **Code:** 200  
  **Content:**
  ```
    {"success": true}
  ```
* **Error Response:**  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Invalid email."
    }
    ```
  OR  
  * **Code:** 404  
  **Content:**
    ```
    {
        "code": 404,
        "description": "Email not found."
    }
    ```
  OR  
  * **Code:** 409  
  **Content:**
    ```
    {
        "code": 409,
        "description": "Subscription already exists."
    }
    ```
  OR  
  * **Code:** 500  
  **Content:**
    ```
    {
        "code": 500,
        "description": "Internal server error."
    }
    ```

**POST /block**
----
  Block updates from an email address.
* **URL Params**  
  None
* **Headers**  
  Content-Type: application/json  
* **Data Params**  
```
  {
    "requestor": <user_email>,
    "target": <user_email>
  }
```
* **Success Response:**  
* **Code:** 200  
  **Content:**
  ```
    {"success": true}
  ```
* **Error Response:**  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Invalid email."
    }
    ```
  OR  
  * **Code:** 404  
  **Content:**
    ```
    {
        "code": 404,
        "description": "Email not found."
    }
    ```
  OR  
  * **Code:** 409  
  **Content:**
    ```
    {
        "code": 409,
        "description": "Blocking already exists."
    }
    ```
  OR  
  * **Code:** 500  
  **Content:**
    ```
    {
        "code": 500,
        "description": "Internal server error."
    }
    ```

**POST /emails-receiving-updates**
----
  Retrieve all email addresses that can receive updates from an email address.
* **URL Params**  
  None
* **Headers**  
  Content-Type: application/json  
* **Data Params**
```
  {
    "sender": <user_email>,
    "text": <user_update>
  }
```
* **Success Response:**  
* **Code:** 200  
  **Content:**
  ```
  {
    "success": true,
    "recipients": [
        <user_email>,
        <user_email>,
        <user_email>
    ]
  }
  ```
* **Error Response:**  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Invalid email."
    }
    ```
  OR  
  * **Code:** 404  
  **Content:**
    ```
    {
        "code": 404,
        "description": "Email not found."
    }
    ```
  OR  
  * **Code:** 500  
  **Content:**
    ```
    {
        "code": 500,
        "description": "Internal server error."
    }
    ```
