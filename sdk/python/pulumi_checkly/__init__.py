# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .alert_channel import *
from .check import *
from .check_group import *
from .dashboard import *
from .maintenance_window import *
from .provider import *
from .snippet import *
from .trigger_check import *
from .trigger_check_group import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_checkly.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_checkly.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "checkly",
  "mod": "index/alertChannel",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/alertChannel:AlertChannel": "AlertChannel"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/check",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/check:Check": "Check"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/checkGroup",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/checkGroup:CheckGroup": "CheckGroup"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/dashboard",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/dashboard:Dashboard": "Dashboard"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/maintenanceWindow",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/maintenanceWindow:MaintenanceWindow": "MaintenanceWindow"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/snippet",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/snippet:Snippet": "Snippet"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/triggerCheck",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/triggerCheck:TriggerCheck": "TriggerCheck"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/triggerCheckGroup",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/triggerCheckGroup:TriggerCheckGroup": "TriggerCheckGroup"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "checkly",
  "token": "pulumi:providers:checkly",
  "fqn": "pulumi_checkly",
  "class": "Provider"
 }
]
"""
)
