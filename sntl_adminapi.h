/**
 * \mainpage Sentinel LDK Administration API
 * \file sntl_adminapi.h Sentinel LDK Administration API declarations
 *
 * Copyright (C) 2023 THALES. All rights reserved.
 *
 * 
 */


#ifndef __SNTL_ADMINAPI_H__
#define __SNTL_ADMINAPI_H__


#  if defined(_MSC_VER) || defined(__WATCOMC__) || defined(__BORLANDC__)
/** unsigned 64 bit */
typedef unsigned __int64 sntl_admin_u64_t;
/** signed 64 bit */
typedef signed __int64 sntl_admin_s64_t;
#  else
/** unsigned 64 bit */
typedef unsigned long long sntl_admin_u64_t;
/** signed 64 bit */
typedef signed long long sntl_admin_s64_t;
#  endif
#  if defined(_MSC_VER)
/** unsigned 32 bit */
typedef unsigned long sntl_admin_u32_t;
/** signed 32 bit   */
typedef signed long sntl_admin_s32_t;
#  else
/** unsigned 32 bit */
typedef unsigned int sntl_admin_u32_t;
/** signed 32 bit   */
typedef signed int sntl_admin_s32_t;
#  endif
/** unsigned 16 bit */
typedef unsigned short sntl_admin_u16_t;
/** signed 16 bit   */
typedef signed short sntl_admin_s16_t;
/** unsigned 8 bit  */
typedef signed char sntl_admin_s8_t;
/** signed 8 bit    */
typedef unsigned char sntl_admin_u8_t;


#if defined(_WIN32) || defined(WIN32) || defined(_MSC_VER) || defined(__BORLANDC__)
#  define SNTL_ADMIN_CALLCONV __stdcall
#else
#  define SNTL_ADMIN_CALLCONV
#endif

#ifdef __cplusplus
extern "C" {
#endif

/**************************************************************************************************/
/**
 * @defgroup sntl_admin_error_codes AdminAPI Status Codes
 *
 * @{
 */

/** Admin API error code definitions */
enum sntl_admin_error_codes
{
    /** Request successfully completed */
    SNTL_ADMIN_STATUS_OK           = 0,

    /** System is out of memory */
    SNTL_ADMIN_INSUF_MEM           = 3,

    /** Admin Context corrupt or contains invalid data */
    SNTL_ADMIN_INVALID_CONTEXT     = 6001,

    /** Specified License Manager could not be found */
    SNTL_ADMIN_LM_NOT_FOUND        = 6002,

    /** Specified License Manager was found, but does not support sntl_admin functions */
    SNTL_ADMIN_LM_TOO_OLD          = 6003,

    /** Errors in request XML - see result for details */
    SNTL_ADMIN_BAD_PARAMETERS      = 6004,

    /** Error initializing the local network interface */
    SNTL_ADMIN_LOCAL_NETWORK_ERR   = 6005,

    /** Error reading file specified by "file://" prefix */
    SNTL_ADMIN_CANNOT_READ_FILE    = 6006,

    /** Error in haspscope */
    SNTL_ADMIN_SCOPE_ERROR         = 6007,

    /** Password is required */
    SNTL_ADMIN_PASSWORD_REQUIRED   = 6008,

    /** Error setting new password */
    SNTL_ADMIN_CANNOT_SET_PASSWORD = 6009,

    /** Error applying an update (see reply for further explanation) */
    SNTL_ADMIN_UPDATE_ERROR        = 6010,

    /** The specified License Manager was found, but the License Manager is configured to block remote access */
    SNTL_ADMIN_LOCAL_ONLY          = 6011,

    /** an XML element tried to set an illegal value */
    SNTL_ADMIN_BAD_VALUE           = 6012,

    /** an XML element tried to set a read-only value */
    SNTL_ADMIN_READ_ONLY           = 6013,

    /** XML element not defined for current request */
    SNTL_ADMIN_ELEMENT_UNDEFINED   = 6014,

    /** An invalid pointer was passed to a sntl_admin_ function */
    SNTL_ADMIN_INVALID_PTR         = 6015,

    /** Standalone AdminAPI cannot access integrated LM */
    SNTL_ADMIN_NO_INTEGRATED_LM    = 6016,

    /** Returned Result is too big - narrow search by using a scope */
    SNTL_ADMIN_RESULT_TOO_BIG      = 6017,

    /** Unable to locate any items matching scope */
    SNTL_ADMIN_SCOPE_RESULTS_EMPTY = 6019,

    /** Invalid Vendor Code passed */
    SNTL_ADMIN_INV_VCODE           = 6022,

    /** Vendor Code not recognized by API */
    SNTL_ADMIN_UNKNOWN_VCODE       = 6034,

    /** Failed to initialize the entropy source */
    SNTL_ADMIN_ENTROPY_SOURCE_FAILED = 6035,

    /** Invalid HTTPS certificate passed in the context creation */
    SNTL_ADMIN_INV_CERT = 6036,

    /** Server certificate verification failed */
    SNTL_ADMIN_CERT_VERIFY_FAILED = 6037,

    /** Connect function is not called */
    SNTL_ADMIN_CONNECT_MISSING     = 6051,

    /** .NET DLL found broken */
    SNTL_ADMIN_NET_DLL_BROKEN      = 6052,

    /** Java JNI DLL loaded in another class loader */
    SNTL_ADMIN_DLL_LOADED_IN_ANOTHER_CLASSLOADER = 6053,

    /** Identity Code is required */
    SNTL_ADMIN_IDENTITY_CODE_REQUIRED   = 6054,

    /** Issued To is required */
    SNTL_ADMIN_ISSUED_TO_REQUIRED   = 6055,

    /** Maximum Number of Auto-registered Machines is required */
    SNTL_ADMIN_MAXIMUM_NUMBER_OF_AUTO_REGISTERED_MACHINES_REQUIRED   = 6056,

    /** Identity Code does not exist */
    SNTL_ADMIN_IDENTITY_CODE_DOES_NOT_EXIST   = 6057,
    
    /** Failed to create identity code */
    SNTL_ADMIN_IDENTITY_CODE_CREATION_FAILED = 6058,

    /** Nothing to unregister */
    SNTL_ADMIN_IDENTITY_NOTHING_TO_UNREGISTER = 6059,

    /** .NET API - Object incorrectly initialized */
    SNTL_ADMIN_INVALID_OBJECT = 6060,

    /** Nothing to update */
    SNTL_ADMIN_IDENTITY_NOTHING_TO_UPDATE = 6061,

    /** Too many identities */
    SNTL_ADMIN_TOO_MANY_IDENTITIES = 6062,

    /** Length of value is too long for the specified field */
    SNTL_ADMIN_LENGTH_TOO_LONG = 6063

};

/**
 * @}
 */



/**
 * @defgroup sntl_admin_general typedefs and macros
 *
 * @{
 */

/** An AdminAPI status code */
typedef enum sntl_admin_error_codes sntl_admin_status_t;

/** Opaque container that holds AdminAPI context */
typedef struct sntl_admin_context_t sntl_admin_context_t;

/**
 * @}
 */


/**
 * @defgroup sntl_adminapi_functions Sentinel LDK AdminAPI functions
 *
 * @{
 */


/**
 * \brief Establishes a session context for subsequent administration requests
 *
 * <b>Sentinel LDK AdminAPI Usage Notes</b><br>
 * This function establishes a context to a Sentinel LDK License Manager
 * This function can't be used to establish a context to Integrated LM of RuntimeLess API.
 *
 * \param admin_context  Unique identifier for a specific License Manager to communicate with
 * \param host           Host name or IP address of the targeted License Manager
 * \param port           Port number of the targeted License Manager (0 is default port 1947)
 * \param password       Pointer to a password string (or NULL)
 *
 * \return            <ul>
 *                      <li>SNTL_ADMIN_STATUS_OK</li> Request was successfully
 *                        completed
 *                      <li>SNTL_ADMIN_INSUF_MEM</li> Out of memory
 *                    </ul>
 *
 * \sa sntl_admin_context_new_scope()
 * \sa sntl_admin_context_delete()
 *
 */

sntl_admin_status_t SNTL_ADMIN_CALLCONV sntl_admin_context_new (
                              sntl_admin_context_t **admin_context,
                              const char            *host,
                              sntl_admin_u16_t       port,
                              const char            *password);

/**
 * \brief Establishes a session context for subsequent administration requests
 *
 * <b>Sentinel LDK AdminAPI Usage Notes</b><br>
 * This function establishes a context to a Sentinel LDK License Manager, or to the
 * Integrated License Manager of the RuntimeLess API
 *
 * \param admin_context  Unique identifier for a specific License Manager to communicate with
 * \param scope          XML specification of the LM to connect to
 *
 *      Connection to a License Manager Service:
 *        <?xml version="1.0" encoding="UTF-8" ?>
 *          <haspscope>
 *	          <host></host> <!--host name or address for LM Service -->
 *	          <port></port> <!--port number for LM service (defaults to 1947) -->
 *            <password></password> <!--If password protection is set in LM Service -->
 *            <certificate></certificate> <!--The CA/Server certificate to verify the server identity. If specified all the communication will be in HTTPS (default is communication in HTTP) -->
 *          </haspscope>
 *
 *      Connection to Integrated LM:
 *        <?xml version="1.0" encoding="UTF-8" ?>
 *          <haspscope>
 *          	<vendor_code></vendor_code> <!-- Required for Integrated LM -->
 *            <host>sntl_integrated_lm</host> <!-- Required for Integrated LM -->
 *          </haspscope>
 *
 *
 * \return            <ul>
 *                      <li>SNTL_ADMIN_STATUS_OK</li> Request was successfully
 *                        completed
 *                      <li>SNTL_ADMIN_INSUF_MEM</li> Out of memory
 *                      <li>SNTL_ADMIN_INV_CERT</li> Invalid HTTPS certificate passed
 *                    </ul>
 *
 * \sa sntl_admin_context_new()
 * \sa sntl_admin_context_delete()
 *
 */

sntl_admin_status_t SNTL_ADMIN_CALLCONV sntl_admin_context_new_scope (
                              sntl_admin_context_t **admin_context,
                              const char            *scope );


/**
 * \brief Frees memory resources utilized by an AdminAPI context.
 *
 * Used to release memory and other resources allocated by creating
 * an AdminAPI context.
 *
 * \param admin_context    Pointer to the context to be released
 *
 * \return                 <ul>
 *                         <li>SNTL_ADMIN_STATUS_OK</li> Request was successfully completed
 *                         </ul>
 *
 * \sa sntl_admin_context_new()
 *
 */
sntl_admin_status_t SNTL_ADMIN_CALLCONV sntl_admin_context_delete (
                              sntl_admin_context_t *admin_context);


/**
 * \brief Gets information from License Manager specified by context
 *
 * Used to retrieve administration or monitoring information from a License Manager
 * specified by an AdminAPI context.
 *
 * \param admin_context    Pointer to the License Manager context to be used
 * \param scope            Definition of the data that is to be searched, in XML
 *                           format. For more information, see the accompanying
 *                           Sentinel LDK AdminAPI help documentation
 * \param format           Definition of the format in which the data is to be
 *                           displayed, in XML format. For more information, see the
 *                           accompanying Sentinel LDK AdminAPI help documentation
 * \param info             Pointer to the information that is retrieved, in XML format, including
 *                           status information. The caller has to free *info after using it.
 *
 * \return                 <ul>
 *                         <li>SNTL_ADMIN_STATUS_OK</li> Request was successfully completed
 *                         <li>SNTL_ADMIN_INSUF_MEM</li> System is out of memory
 *                         <li>SNTL_ADMIN_INVALID_CONTEXT</li> Admin Context corrupt or contains
 *                           invalid data
 *                         <li>SNTL_ADMIN_SCOPE_ERROR</li>Error in haspscope
 *                         <li>SNTL_ADMIN_LM_NOT_FOUND</li> Specified License Manager could
 *                           not be found
 *                         <li>SNTL_ADMIN_LM_TOO_OLD</li> Specified License Manager was found,
 *                           but does not support sntl_admin functions
 *                         <li>SNTL_ADMIN_BAD_PARAMETERS</li> Errors in request XML - see
 *                           resulting info XML for details
 *                         <li>SNTL_ADMIN_LOCAL_NETWORK_ERR</li> Error initializing the local
 *                           network interface
 *                         <li>SNTL_ADMIN_PASSWORD_REQUIRED</li> Password is required
 *                         <li>SNTL_ADMIN_INVALID_PTR</li> "format" or "scope" parameter is NULL
 *                         </ul>
 *
 * \sa sntl_admin_free()
 * \sa sntl_admin_context_new()
 *
 */

sntl_admin_status_t SNTL_ADMIN_CALLCONV sntl_admin_get (
                              sntl_admin_context_t  *admin_context,
                              const char            *scope,
                              const char            *format,
                              char                 **info);

/**
 * \brief Sets configuration items forLicense Manager specified by context
 *
 * Used to set configuration options or apply actions to a License Manager specified by
 * an AdminAPI context.
 *
 * \param admin_context    Pointer to the License Manager context to be used
 * \param input            Definition of the actions/settings to be done, in XML format.
 *                           For more information, see the accompanying Sentinel LDK AdminAPI
 *                           help documentation
 * \param status           Pointer to the status information for the applied actions/settings,
 *                           in XML format. The caller has to free *status after using it.
 *
 * \return                 <ul>
 *                         <li>SNTL_ADMIN_STATUS_OK</li> Request was successfully completed
 *                         <li>SNTL_ADMIN_INSUF_MEM</li> System is out of memory
 *                         <li>SNTL_ADMIN_INVALID_CONTEXT</li> Admin Context corrupt or contains
 *                           invalid data
 *                         <li>SNTL_ADMIN_LM_NOT_FOUND</li> Specified License Manager could
 *                           not be found
 *                         <li>SNTL_ADMIN_LM_TOO_OLD</li> Specified License Manager was found,
 *                           but does not support sntl_admin functions
 *                         <li>SNTL_ADMIN_BAD_PARAMETERS</li> Errors in request XML - see
 *                           resulting info XML for details
 *                         <li>SNTL_ADMIN_LOCAL_NETWORK_ERR</li> Error initializing the local
 *                           network interface
 *                         <li>SNTL_ADMIN_CANNOT_READ_FILE</li>Error reading file specified by
 *                           "file://" prefix
 *                         <li>SNTL_ADMIN_PASSWORD_REQUIRED</li> Password is required
 *                         <li>SNTL_ADMIN_CANNOT_SET_PASSWORD</li> Error setting new password
 *                         <li>SNTL_ADMIN_UPDATE_ERROR</li> Error applying an update (see reply
 *                           XML for further explanation)
 *                         <li>SNTL_ADMIN_BAD_VALUE</li> An XML element tried to set an illegal value
 *                         <li>SNTL_ADMIN_INVALID_PTR</li> "input" parameter is NULL
 *                         </ul>
 *
 * \sa sntl_admin_free()
 * \sa sntl_admin_context_new()
 *
 */

sntl_admin_status_t SNTL_ADMIN_CALLCONV sntl_admin_set (
                              sntl_admin_context_t  *admin_context,
                              const char            *input,
                              char                 **status );


/**
 * \brief Frees memory resources utilized by other AdminAPI functions.
 *
 * Used to release memory resources allocated to store retrieved data from AdminAPI calls.
 *
 * \param info         Pointer to the memory resources to be released
 * \return             <ul>
 *                     <li>SNTL_ADMIN_STATUS_OK</li> Request was successfully completed
 *                     </ul>
 *
 * \sa sntl_admin_get()
 * \sa sntl_admin_set()
 *
 */

void SNTL_ADMIN_CALLCONV sntl_admin_free ( char *info );


/**
 * @}
 */

/**************************************************************************************************/

#ifdef __cplusplus
} // extern "C"
#endif

#endif
