package undoconfig

import (
	"github.com/git-town/git-town/v20/internal/config/configdomain"
	"github.com/git-town/git-town/v20/internal/gohacks/mapstools"
	"github.com/git-town/git-town/v20/internal/vm/opcodes"
	"github.com/git-town/git-town/v20/internal/vm/program"
)

// ConfigDiffs describes the changes made to the local and global Git configuration.
type ConfigDiffs struct {
	Global ConfigDiff
	Local  ConfigDiff
}

func NewConfigDiffs(before, after ConfigSnapshot) ConfigDiffs {
	return ConfigDiffs{
		Global: SingleCacheDiff(before.Global, after.Global),
		Local:  SingleCacheDiff(before.Local, after.Local),
	}
}

func (self ConfigDiffs) UndoProgram() program.Program {
	result := program.Program{}
	for _, key := range self.Global.Added {
		result.Add(&opcodes.ConfigRemove{
			Key:   key,
			Scope: configdomain.ConfigScopeGlobal,
		})
	}
	for key, value := range mapstools.SortedKeyValues(self.Global.Removed) {
		result.Add(&opcodes.ConfigSet{
			Key:   key,
			Scope: configdomain.ConfigScopeGlobal,
			Value: value,
		})
	}
	for key, value := range mapstools.SortedKeyValues(self.Global.Changed) {
		result.Add(&opcodes.ConfigSet{
			Key:   key,
			Scope: configdomain.ConfigScopeGlobal,
			Value: value.Before,
		})
	}
	for _, key := range self.Local.Added {
		result.Add(&opcodes.ConfigRemove{
			Key:   key,
			Scope: configdomain.ConfigScopeLocal,
		})
	}
	for key, value := range mapstools.SortedKeyValues(self.Local.Removed) {
		result.Add(&opcodes.ConfigSet{
			Key:   key,
			Scope: configdomain.ConfigScopeLocal,
			Value: value,
		})
	}
	for key, value := range mapstools.SortedKeyValues(self.Local.Changed) {
		result.Add(&opcodes.ConfigSet{
			Key:   key,
			Scope: configdomain.ConfigScopeLocal,
			Value: value.Before,
		})
	}
	return result
}
