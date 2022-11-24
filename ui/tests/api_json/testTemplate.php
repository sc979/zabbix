<?php
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


require_once dirname(__FILE__).'/../include/CAPITest.php';

/**
 * @backup hstgrp
 * @backup hosts
 */
class testTemplate extends CAPITest {

	public static function dataProviderCreate() {
		return [
			[
				'request' => [
					'host' => 'test-template-01',
					'groups' => ['groupid' => 1]
				]
			],
			[
				'request' => [
					'host' => 'test-template-02',
					'groups' => []
				],
				'expected_error' => 'Invalid parameter "/1/groups": cannot be empty.'
			],
			[
				'request' => [
					'host' => 'test-template-03',
					'groups' => [[]]
				],
				'expected_error' => 'Invalid parameter "/1/groups/1": the parameter "groupid" is missing.'
			],
			[
				'request' => [
					'host' => 'test-template-04',
					'groups' => ['groupid' => 9999]
				],
				'expected_error' => 'No permissions to referred object or it does not exist!'
			],
			[
				'request' => [
					'host' => 'test-template-05',
					'groups' => ['groupid' => 1]
				],
				'expected_error' => 'No permissions to referred object or it does not exist!',
				'user' => ['user' => 'zabbix-admin', 'password' => 'zabbix']
			],
			[
				'request' => [
					[
						'host' => 'test-template-06',
						'groups' => ['groupid' => 1]
					],
					[
						'host' => 'test-template-06',
						'groups' => ['groupid' => 1]
					]
				],
				'expected_error' => 'Invalid parameter "/2": value (host)=(test-template-06) already exists.'
			],
			[
				'request' => [
					[
						'host' => 'test-template-07',
						'name' => 'test-template-07',
						'groups' => ['groupid' => 1]
					],
					[
						'host' => 'test-template-08',
						'name' => 'test-template-07',
						'groups' => ['groupid' => 1]
					]
				],
				'expected_error' => 'Invalid parameter "/2": value (name)=(test-template-07) already exists.'
			],
			[
				'request' => [
					[
						'host' => 'test-template-09',
						'groups' => ['groupid' => 1]
					]
				]
			],
			// The next two test cases depends on the previous one.
			[
				'request' => [
					[
						'host' => 'test-template-09',
						'groups' => ['groupid' => 1]
					]
				],
				'expected_error' => 'Template with host name "test-template-09" already exists.'
			],
			[
				'request' => [
					[
						'host' => 'test-template-10',
						'name' => 'test-template-09',
						'groups' => ['groupid' => 1]
					]
				],
				'expected_error' => 'Template with visible name "test-template-09" already exists.'
			],
			// The next two test cases depends on the existing host "Zabbix server".
			[
				'request' => [
					[
						'host' => 'Zabbix server',
						'groups' => ['groupid' => 1]
					]
				],
				'expected_error' => 'Host with host name "Zabbix server" already exists.'
			],
			[
				'request' => [
					[
						'host' => 'test-template-11',
						'name' => 'Zabbix server',
						'groups' => ['groupid' => 1]
					]
				],
				'expected_error' => 'Host with visible name "Zabbix server" already exists.'
			],
			[
				'request' => [
					[
						'host' => 'test-template-12',
						'groups' => ['groupid' => 1],
						'templates' => ['templateid' => 10047 /* "Zabbix server health" */]
					]
				],
				'expected_error' => 'Invalid parameter "/1": unexpected parameter "templates".'
			]
		];
	}

	/**
	 * @dataProvider dataProviderCreate
	 */
	public function testTemplate_Create(array $request, string $expected_error = null, array $user = null) {
		if ($user !== null) {
			$this->authorize($user['user'], $user['password']);
		}
		$this->call('template.create', $request, $expected_error);
	}

	public function testTemplate_CreateUUID() {
		$result = $this->call('template.create', [
			[
				'host' => 'test-template-uuid-01',
				'groups' => ['groupid' => 1]
			],
			[
				'host' => 'test-template-uuid-02',
				'groups' => ['groupid' => 1]
			]
		])['result'];

		$this->assertArrayHasKey('templateids', $result);

		$db_templates = $this->call('template.get', [
			'output' => ['uuid'],
			'templateids' => $result['templateids']
		])['result'];

		$this->assertTrue(preg_match('/[0-9a-f]{32}/', $db_templates[0]['uuid']) === 1);
		$this->assertTrue(preg_match('/[0-9a-f]{32}/', $db_templates[1]['uuid']) === 1);
		$this->assertNotSame($db_templates[0]['uuid'], $db_templates[1]['uuid']);
	}
}
