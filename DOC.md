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

**POST /users**
----
  Create a new user.
* **URL Params**  
  None
* **Headers**  
  Content-Type: application/json  
* **Data Params**  
```
  {
    "email": <user_email>
  }
```
* **Success Response:**  
* **Code:** 200  
  **Content:**
  ```
    {"success": true}
  ```

**POST /friends**
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
        "description": "Invalid JSON request"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Invalid email"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "The number of emails must be 2"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "The emails are the same"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "user not found"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "blocking relationship exists"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "friend connection exists"
    }
    ```
  OR  
  * **Code:** 500  
  **Content:**
    ```
    {
        "code": 500,
        "description": "Internal Server Error"
    }
    ```

**POST /friends/list**
----
  Get the friends list for an email address.
* **URL Params**  
  None
* **Headers**  
  Content-Type: application/json  
* **Data Params**  
```
  {
    "email": <user_email>
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
        "description": "Invalid JSON request"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Missing email field"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Invalid email"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "user not found"
    }
    ```
  OR  
  * **Code:** 500  
  **Content:**
    ```
    {
        "code": 500,
        "description": "Internal Server Error"
    }
    ```

**POST /friends/common**
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
        "description": "Invalid JSON request"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Invalid email"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "The number of emails must be 2"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "The emails are the same"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "user not found"
    }
    ```
  OR  
  * **Code:** 500  
  **Content:**
    ```
    {
        "code": 500,
        "description": "Internal Server Error"
    }
    ```

**POST /friends/subscription**
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
        "description": "Invalid JSON request"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Missing requestor field"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Missing target field"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Invalid email"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Requestor and target are the same"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "user not found"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "subscription exists"
    }
    ```
  OR  
  * **Code:** 500  
  **Content:**
    ```
    {
        "code": 500,
        "description": "Internal Server Error"
    }
    ```

**POST /friends/block**
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
        "description": "Invalid JSON request"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Missing requestor field"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Missing target field"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Invalid email"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Requestor and target are the same"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "user not found"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "blocking relationship exists"
    }
    ```
  OR  
  * **Code:** 500  
  **Content:**
    ```
    {
        "code": 500,
        "description": "Internal Server Error"
    }
    ```

**POST /friends/recipients**
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
        "description": "Invalid JSON request"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Missing sender field"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Missing text field"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "Invalid email"
    }
    ```
  OR  
  * **Code:** 400  
  **Content:**
    ```
    {
        "code": 400,
        "description": "user not found"
    }
    ```
  OR  
  * **Code:** 500  
  **Content:**
    ```
    {
        "code": 500,
        "description": "Internal Server Error"
    }
    ```
