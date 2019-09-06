/*
** Zabbix
** Copyright (C) 2001-2019 Zabbix SIA
**
** This program is free software; you can redistribute it and/or modify
** it under the terms of the GNU General Public License as published by
** the Free Software Foundation; either version 2 of the License, or
** (at your option) any later version.
**
** This program is distributed in the hope that it will be useful,
** but WITHOUT ANY WARRANTY; without even the implied warranty of
** MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
** GNU General Public License for more details.
**
** You should have received a copy of the GNU General Public License
** along with this program; if not, write to the Free Software
** Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
**/

package zbxlib

/* cspell:disable */

/*
#cgo CFLAGS: -I${SRCDIR}/../../../../../include

#cgo LDFLAGS: -Wl,--start-group
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/zabbix_agent/logfiles/libzbxlogfiles.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxcomms/libzbxcomms.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxcommon/libzbxcommon.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxcrypto/libzbxcrypto.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxsys/libzbxsys.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxnix/libzbxnix.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxconf/libzbxconf.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxhttp/libzbxhttp.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxcompress/libzbxcompress.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxregexp/libzbxregexp.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxsysinfo/libzbxagentsysinfo.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxsysinfo/common/libcommonsysinfo.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxsysinfo/simple/libsimplesysinfo.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxsysinfo/linux/libspechostnamesysinfo.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxsysinfo/linux/libspecsysinfo.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxexec/libzbxexec.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxalgo/libzbxalgo.a
#cgo LDFLAGS: ${SRCDIR}/../../../../../src/libs/zbxjson/libzbxjson.a
#cgo LDFLAGS: -lz -lpcre -lresolv -lcurl
#cgo LDFLAGS: -Wl,--end-group

#include "common.h"
#include "sysinfo.h"
#include "comms.h"
#include "../src/zabbix_agent/metrics.h"
#include "../src/zabbix_agent/logfiles/logfiles.h"

typedef ZBX_ACTIVE_METRIC* ZBX_ACTIVE_METRIC_LP;
typedef zbx_vector_ptr_t * zbx_vector_ptr_lp_t;

int CONFIG_MAX_LINES_PER_SECOND = 20;
char *CONFIG_HOSTNAME = NULL;
int	CONFIG_UNSAFE_USER_PARAMETERS= 0;
int	CONFIG_ENABLE_REMOTE_COMMANDS= 0;
int	CONFIG_LOG_REMOTE_COMMANDS= 0;
char	*CONFIG_SOURCE_IP= NULL;

unsigned int	configured_tls_connect_mode = ZBX_TCP_SEC_UNENCRYPTED;
unsigned int	configured_tls_accept_modes = ZBX_TCP_SEC_UNENCRYPTED;

char *CONFIG_TLS_CONNECT= NULL;
char *CONFIG_TLS_ACCEPT	= NULL;
char *CONFIG_TLS_CA_FILE = NULL;
char *CONFIG_TLS_CRL_FILE = NULL;
char *CONFIG_TLS_SERVER_CERT_ISSUER	= NULL;
char *CONFIG_TLS_SERVER_CERT_SUBJECT = NULL;
char *CONFIG_TLS_CERT_FILE = NULL;
char *CONFIG_TLS_KEY_FILE = NULL;
char *CONFIG_TLS_PSK_IDENTITY = NULL;
char *CONFIG_TLS_PSK_FILE = NULL;

int	CONFIG_PASSIVE_FORKS = 0;
int	CONFIG_ACTIVE_FORKS = 0;

const char	*progname = NULL;
const char	title_message[] = "agent";
const char	syslog_app_name[] = "agent";
const char	*usage_message[] = {};
unsigned char	program_type	= 0x80;
const char	*help_message[] = {};

ZBX_METRIC	parameters_agent[] = {NULL};
ZBX_METRIC	parameters_specific[] = {NULL};

void zbx_on_exit(int ret)
{
}

int	zbx_procstat_collector_started(void)
{
	return FAIL;
}

int	zbx_procstat_get_util(const char *procname, const char *username, const char *cmdline, zbx_uint64_t flags,
		int period, int type, double *value, char **errmsg)
{
	return FAIL;
}

int	get_cpustat(AGENT_RESULT *result, int cpu_num, int state, int mode)
{
	return SYSINFO_RET_FAIL;
}
*/
import "C"

const (
	ItemStateNormal       = 0
	ItemStateNotsupported = 1
)

const (
	Succeed = 0
	Fail    = -1
)
