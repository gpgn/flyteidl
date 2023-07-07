# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from flyteadmin.models.secret_mount_env_var import SecretMountEnvVar  # noqa: F401,E501
from flyteadmin.models.secret_mount_file import SecretMountFile  # noqa: F401,E501
from flyteadmin.models.secret_mount_type import SecretMountType  # noqa: F401,E501


class CoreSecret(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'group': 'str',
        'group_version': 'str',
        'key': 'str',
        'mount_requirement': 'SecretMountType',
        'env_var': 'SecretMountEnvVar',
        'file': 'SecretMountFile'
    }

    attribute_map = {
        'group': 'group',
        'group_version': 'group_version',
        'key': 'key',
        'mount_requirement': 'mount_requirement',
        'env_var': 'env_var',
        'file': 'file'
    }

    def __init__(self, group=None, group_version=None, key=None, mount_requirement=None, env_var=None, file=None):  # noqa: E501
        """CoreSecret - a model defined in Swagger"""  # noqa: E501

        self._group = None
        self._group_version = None
        self._key = None
        self._mount_requirement = None
        self._env_var = None
        self._file = None
        self.discriminator = None

        if group is not None:
            self.group = group
        if group_version is not None:
            self.group_version = group_version
        if key is not None:
            self.key = key
        if mount_requirement is not None:
            self.mount_requirement = mount_requirement
        if env_var is not None:
            self.env_var = env_var
        if file is not None:
            self.file = file

    @property
    def group(self):
        """Gets the group of this CoreSecret.  # noqa: E501


        :return: The group of this CoreSecret.  # noqa: E501
        :rtype: str
        """
        return self._group

    @group.setter
    def group(self, group):
        """Sets the group of this CoreSecret.


        :param group: The group of this CoreSecret.  # noqa: E501
        :type: str
        """

        self._group = group

    @property
    def group_version(self):
        """Gets the group_version of this CoreSecret.  # noqa: E501


        :return: The group_version of this CoreSecret.  # noqa: E501
        :rtype: str
        """
        return self._group_version

    @group_version.setter
    def group_version(self, group_version):
        """Sets the group_version of this CoreSecret.


        :param group_version: The group_version of this CoreSecret.  # noqa: E501
        :type: str
        """

        self._group_version = group_version

    @property
    def key(self):
        """Gets the key of this CoreSecret.  # noqa: E501


        :return: The key of this CoreSecret.  # noqa: E501
        :rtype: str
        """
        return self._key

    @key.setter
    def key(self, key):
        """Sets the key of this CoreSecret.


        :param key: The key of this CoreSecret.  # noqa: E501
        :type: str
        """

        self._key = key

    @property
    def mount_requirement(self):
        """Gets the mount_requirement of this CoreSecret.  # noqa: E501


        :return: The mount_requirement of this CoreSecret.  # noqa: E501
        :rtype: SecretMountType
        """
        return self._mount_requirement

    @mount_requirement.setter
    def mount_requirement(self, mount_requirement):
        """Sets the mount_requirement of this CoreSecret.


        :param mount_requirement: The mount_requirement of this CoreSecret.  # noqa: E501
        :type: SecretMountType
        """

        self._mount_requirement = mount_requirement

    @property
    def env_var(self):
        """Gets the env_var of this CoreSecret.  # noqa: E501


        :return: The env_var of this CoreSecret.  # noqa: E501
        :rtype: SecretMountEnvVar
        """
        return self._env_var

    @env_var.setter
    def env_var(self, env_var):
        """Sets the env_var of this CoreSecret.


        :param env_var: The env_var of this CoreSecret.  # noqa: E501
        :type: SecretMountEnvVar
        """

        self._env_var = env_var

    @property
    def file(self):
        """Gets the file of this CoreSecret.  # noqa: E501


        :return: The file of this CoreSecret.  # noqa: E501
        :rtype: SecretMountFile
        """
        return self._file

    @file.setter
    def file(self, file):
        """Sets the file of this CoreSecret.


        :param file: The file of this CoreSecret.  # noqa: E501
        :type: SecretMountFile
        """

        self._file = file

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(CoreSecret, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, CoreSecret):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
