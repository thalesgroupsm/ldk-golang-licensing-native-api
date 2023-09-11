
#ifndef __HASP_API_INT_H__
#define __HASP_API_INT_H__
#include "hasp_api.h"
#ifdef __cplusplus
extern "C" {
#endif

/*! \brief Cleanup memory before terminating.
 *
 * This function is used to release all the resources allocated by the API before
 * the process termination.
 *
 * Resources are anyway automatically released by the operating system when the application
 * terminates, but when using a memory leak detector, you can use hasp_cleanup() to force
 * the deallocation of the memory used by the HASP API before the application termination,
 * avoiding to have it reported as leaked.
 *
 * Note that after calling hasp_cleanup() no other HASP functions should be called.
 *
 * This function never fails, and it always returns HASP_STATUS_OK.
 */
hasp_status_t HASP_CALLCONV hasp_cleanup(void);
#ifdef __cplusplus
}
#endif
#endif