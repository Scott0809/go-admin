// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package language

import "strings"

var en = LangSet{
	"managers":         "Managers",
	"name":             "Name",
	"nickname":         "Nickname",
	"role":             "Role",
	"createdat":        "createdAt",
	"updatedat":        "updatedAt",
	"path":             "path",
	"new":              "New",
	"action":           "Action",
	"toggle dropdown":  "Toggle Dropdown",
	"delete":           "Delete",
	"refresh":          "Refresh",
	"expand":           "Expand",
	"collapse":         "Collapse",
	"back":             "Back",
	"reset":            "Reset",
	"save":             "Save",
	"edit":             "Edit",
	"operation":        "Operation",
	"method":           "Method",
	"input":            "input",
	"online":           "Online",
	"setting":          "Setting",
	"sign out":         "Sign out",
	"all":              "All",
	"confirm password": "Confirm Password",
	"search":           "Search",
	"remove":           "Remove",

	"are you sure to delete":               "Are you sure to delete",
	"yes":                                  "yes",
	"got it":                               "got it",
	"cancel":                               "cancel",
	"refresh succeeded":                    "Refresh succeeded",
	"reload succeeded":                     "Reload succeeded",
	"all method if empty":                  "All method if empty",
	"password does not match":              "Password does not match",
	"should be unique":                     "Should be unique",
	"slug exists":                          "Slug exists",
	"no corresponding options?":            "No corresponding options?",
	"create here.":                         "Create here.",
	"use for login":                        "Use for login",
	"use to display":                       "Use to display",
	"a path a line, without global prefix": "A path a line",
	"slug or http_path or name should not be empty": "slug or http_path or name should not be empty",
	"no roles":          "no roles",
	"fixed the sidebar": "Fixed the sidebar",
	"enter fullscreen":  "Enter fullscreen",
	"exit fullscreen":   "Exit fullscreen",

	"permission manage": "Permission Manage",
	"menus manage":      "Menus Manage",
	"roles manage":      "Roles manage",
	"operation log":     "Operation log",

	"continue editing":  "Continue editing",
	"continue creating": "Continue creating",

	"browse":     "Browse",
	"avatar":     "Avatar",
	"password":   "Password",
	"username":   "Username",
	"slug":       "Slug",
	"permission": "Permission",
	"userid":     "UserID",
	"content":    "Content",
	"parent":     "Parent",
	"icon":       "Icon",
	"uri":        "Uri",

	"detail": "Detail",

	"admin":     "Admin",
	"users":     "Users",
	"roles":     "Roles",
	"menu":      "Menu",
	"dashboard": "Dashboard",
	"home":      "Home",

	"second":  "second",
	"seconds": "seconds",
	"minute":  "minute",
	"minutes": "minutes",
	"hour":    "hour",
	"hours":   "hours",
	"day":     "day",
	"days":    "days",
	"week":    "week",
	"weeks":   "weeks",
	"month":   "month",
	"months":  "months",
	"year":    "year",
	"years":   "years",

	"config.domain":          "Website Domain",
	"config.language":        "Website Language",
	"config.url prefix":      "URL Prefix",
	"config.theme":           "Theme",
	"config.title":           "Title",
	"config.index url":       "Home URL",
	"config.login url":       "Login URL",
	"config.env":             "Env",
	"config.color scheme":    "Color Scheme",
	"config.cdn url":         "CDN Asset URL",
	"config.login title":     "Login Title",
	"config.auth user table": "Auth User Table",
	"config.extra":           "Extra Configuration",
	"config.store":           "File Store Setting",
	"config.databases":       "Database Setting",

	"config.general":      "General",
	"config.log":          "Log",
	"config.site setting": "Site Settings",
	"config.custom":       "Customize",
	"config.debug":        "Debug Mode",
	"config.site off":     "Site Offline",
	"config.true":         "On",
	"config.false":        "Off",

	"config.logo":                        "Logo",
	"config.mini logo":                   "Mini Logo",
	"config.session life time":           "Session Life Time",
	"config.custom head html":            "Head HTML",
	"config.custom foot html":            "Foot HTML",
	"config.custom 404 html":             "404 Page",
	"config.custom 403 html":             "403 Page",
	"config.custom 500 html":             "500 Page",
	"config.hide config center entrance": "Hide Config Button",
	"config.hide app info entrance":      "Hide App Info Button",
	"config.hide tool entrance":          "Hide Tool Button",
	"config.footer info":                 "Footer Info",
	"config.login logo":                  "Login Logo",
	"config.no limit login ip":           "No Limit Login Multi IPs",
	"config.animation type":              "Animation Type",
	"config.animation duration":          "Animation Duration(s)",
	"config.animation delay":             "Animation Delay(s)",
	"config.file upload engine":          "File Upload Engine",

	"config.logger rotate":             "Log Rotate Settings",
	"config.logger rotate max size":    "Max Size（m）",
	"config.logger rotate max backups": "Max Buckups",
	"config.logger rotate max age":     "Max Age（day）",
	"config.logger rotate compress":    "Compress",

	"config.info log path":         "Info Log File Path",
	"config.error log path":        "Error Log File Path",
	"config.access log path":       "Access Log File Path",
	"config.info log off":          "Info Log Off",
	"config.error log off":         "Error Log Off",
	"config.access log off":        "Access Log Off",
	"config.access assets log off": "Access Assets Log Off",
	"config.sql log on":            "Open SQL Log",
	"config.log level":             "Level",

	"config.logger rotate encoder":                "Log Encoder Settings",
	"config.logger rotate encoder time key":       "Time Key",
	"config.logger rotate encoder level key":      "Level Key",
	"config.logger rotate encoder name key":       "Name Key",
	"config.logger rotate encoder caller key":     "Caller Key",
	"config.logger rotate encoder message key":    "Message Key",
	"config.logger rotate encoder stacktrace key": "Stacktrace Key",
	"config.logger rotate encoder level":          "Level Encoder",
	"config.logger rotate encoder time":           "Time Encoder",
	"config.logger rotate encoder duration":       "Duration Encoder",
	"config.logger rotate encoder caller":         "Caller Encoder",
	"config.logger rotate encoder encoding":       "Output Format",

	"config.capital":        "Capital",
	"config.capitalcolor":   "Capital with color",
	"config.lowercase":      "Lowercase",
	"config.lowercasecolor": "Lowercase with color",

	"config.seconds":     "Seconds",
	"config.nanosecond":  "Nanosecond",
	"config.microsecond": "Microsecond",
	"config.millisecond": "Millisecond",

	"config.full path":  "Full path",
	"config.short path": "Short path",

	"config.do not modify when you have not set up all assets": "Do not modify when you have not set up all assets",
	"config.it will work when theme is adminlte":               "It will work when theme is adminlte",

	"config.language." + CN:                  "Chinese",
	"config.language." + EN:                  "English",
	"config.language." + JP:                  "Japanese",
	"config.language." + strings.ToLower(TC): "Traditional Chinese",

	"config.modify site config":         "Site Configuration Modification",
	"config.modify site config success": "modified success",
	"config.modify site config fail":    "modified failed",

	"system.system info":     "System And Application Info",
	"system.application":     "Application Info",
	"system.application run": "Applications Running Info",
	"system.system":          "System Info",

	"system.process_id":                           "Process ID",
	"system.golang_version":                       "Golang Version",
	"system.server_uptime":                        "Server Uptime",
	"system.current_goroutine":                    "Current Goroutines",
	"system.current_memory_usage":                 "Current Memory Usage",
	"system.total_memory_allocated":               "Total Memory Allocated",
	"system.memory_obtained":                      "Memory Obtained",
	"system.pointer_lookup_times":                 "Pointer Lookup Times",
	"system.memory_allocate_times":                "Memory Allocate Times",
	"system.memory_free_times":                    "Memory Free Times",
	"system.current_heap_usage":                   "Current Heap Usage",
	"system.heap_memory_obtained":                 "Heap Memory Obtained",
	"system.heap_memory_idle":                     "Heap Memory Idle",
	"system.heap_memory_in_use":                   "Heap Memory In Use",
	"system.heap_memory_released":                 "Heap Memory Released",
	"system.heap_objects":                         "Heap Objects",
	"system.bootstrap_stack_usage":                "Bootstrap Stack Usage",
	"system.stack_memory_obtained":                "Stack Memory Obtained",
	"system.mspan_structures_usage":               "MSpan Structures Usage",
	"system.mspan_structures_obtained":            "MSpan Structures Obtained",
	"system.mcache_structures_usage":              "MCache Structures Usage",
	"system.mcache_structures_obtained":           "MCache Structures Obtained",
	"system.profiling_bucket_hash_table_obtained": "Profiling Bucket Hash Table Obtained",
	"system.gc_metadata_obtained":                 "GC Metadata Obtained",
	"system.other_system_allocation_obtained":     "Other System Allocation Obtained",
	"system.next_gc_recycle":                      "Next GC Recycle",
	"system.last_gc_time":                         "Since Last GC Time",
	"system.total_gc_time":                        "Total GC Pause",
	"system.total_gc_pause":                       "Total GC Pause",
	"system.last_gc_pause":                        "Last GC Pause",
	"system.gc_times":                             "GC Times",

	"system.cpu_logical_core": "CPU Logical Core",
	"system.cpu_core":         "CPU Physical Core",
	"system.os_platform":      "OS Platform",
	"system.os_family":        "OS Family",
	"system.os_version":       "OS Version",
	"system.load1":            "Load1",
	"system.load5":            "Load5",
	"system.load15":           "Load15",
	"system.mem_total":        "Total Memory",
	"system.mem_available":    "Available Memory",
	"system.mem_used":         "Used Memory",

	"system.app_name":         "App Name",
	"system.go_admin_version": "App Version",
	"system.theme_name":       "Theme",
	"system.theme_version":    "Theme Version",

	"tool.tool":                   "Tool",
	"tool.table":                  "Table",
	"tool.connection":             "Connection",
	"tool.package":                "Package",
	"tool.output":                 "Output Path",
	"tool.output path is empty":   "Output path is empty",
	"tool.field":                  "Field",
	"tool.title":                  "Title",
	"tool.field name":             "Name",
	"tool.db type":                "Database Type",
	"tool.form type":              "Form Type",
	"tool.generate table model":   "Generate Table Model",
	"tool.primarykey":             "Primary Key",
	"tool.field filterable":       "Filterable",
	"tool.field sortable":         "Sortable",
	"tool.yes":                    "Yes",
	"tool.no":                     "No",
	"tool.hide":                   "Hide",
	"tool.show":                   "Show",
	"tool.generate success":       "Generate Success",
	"tool.hide filter area":       "Hide Filter Area",
	"tool.use absolute path":      "Use absolute path",
	"tool.display":                "Display",
	"tool.basic info":             "Basic",
	"tool.table info":             "Table",
	"tool.form info":              "Form",
	"tool.field editable":         "Editable",
	"tool.field can add":          "Can Add",
	"tool.filter area":            "Filter Area",
	"tool.new button":             "New Button",
	"tool.export button":          "Export Button",
	"tool.edit button":            "Edit Button",
	"tool.delete button":          "Delete Button",
	"tool.detail button":          "Detail Button",
	"tool.filter button":          "Filter Button",
	"tool.row selector":           "Row Selector",
	"tool.pagination":             "Pagination",
	"tool.query info":             "Query Info",
	"tool.filter form layout":     "Filter Form Layout",
	"tool.continue edit checkbox": "Continue Edit Checkbox",
	"tool.continue new checkbox":  "Continue New Checkbox",
	"tool.reset button":           "Reset Button",
	"tool.back button":            "Back Button",

	"tool.generate table model success": "generate success",
	"tool.generate table model fail":    "generate fail",

	"plugin.plugin":         "Plugin",
	"plugin.plugin detail":  "Plugin Detail",
	"plugin.introduction":   "Introduction",
	"plugin.website":        "Website",
	"plugin.version":        "Version",
	"plugin.created at":     "Created At",
	"plugin.updated at":     "Updated At",
	"plugin.provided by %s": "Provided by %s",
	"plugin.upgrade":        "Upgrade",
	"plugin.install":        "Install",
	"plugin.info":           "Detail",
}
