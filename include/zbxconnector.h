/*
** Zabbix
** Copyright (C) 2001-2022 Zabbix SIA
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

#ifndef ZABBIX_CONNECTOR_H
#define ZABBIX_CONNECTOR_H

#include "zbxtypes.h"
#include "zbxipcservice.h"

#define ZBX_IPC_SERVICE_CONNECTOR		"connector"

#define ZBX_IPC_CONNECTOR_WORKER		1
#define ZBX_IPC_CONNECTOR_REQUEST		2

typedef struct
{
	zbx_uint64_t	objectid;
	zbx_timespec_t	ts;
	char		*str;
}
zbx_connector_object_t;

ZBX_PTR_VECTOR_DECL(connector_object_ptr, zbx_connector_object_t *)

void	zbx_connector_serialize_object(unsigned char **data, size_t *data_alloc, size_t *data_offset,
		const zbx_connector_object_t *connector_object);
void	zbx_connector_deserialize_object(const unsigned char *data, zbx_uint32_t size,
		zbx_vector_connector_object_ptr_t *connector_objects);
void	zbx_connector_object_free(zbx_connector_object_t *connector_object);

void	zbx_connector_send(zbx_uint32_t code, unsigned char *data, zbx_uint32_t size);
#endif /* ZABBIX_AVAILABILITY_H */
