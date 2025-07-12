# API Specifications

## 1. LC Issuance

*   **Endpoint:** `/lcs`
*   **Method:** `POST`
*   **Request Body:**
    ```json
    {
      "importer": "importer_id",
      "exporter": "exporter_id",
      "amount": 100000,
      "expiryDate": "2023-12-31"
    }
    ```
*   **Response Body:**
    ```json
    {
      "lcId": "lc_id"
    }
    ```

## 2. Document Submission

*   **Endpoint:** `/lcs/{lcId}/documents`
*   **Method:** `POST`
*   **Request Body:**
    ```json
    {
      "documents": [
        {
          "name": "invoice.pdf",
          "hash": "document_hash"
        },
        {
          "name": "bill_of_lading.pdf",
          "hash": "document_hash"
        }
      ]
    }
    ```
*   **Response Body:**
    ```json
    {
      "lcId": "lc_id"
    }
    ```

## 3. Payment Settlement

*   **Endpoint:** `/lcs/{lcId}/settlement`
*   **Method:** `POST`
*   **Response Body:**
    ```json
    {
      "lcId": "lc_id"
    }
    ```
