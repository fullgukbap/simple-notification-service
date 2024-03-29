// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"notification-service/ent/entitytype"
	"notification-service/ent/friendship"
	"notification-service/ent/friendshipstatus"
	"notification-service/ent/notification"
	"notification-service/ent/notificationchange"
	"notification-service/ent/notificationobject"
	"notification-service/ent/schema"
	"notification-service/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	entitytypeMixin := schema.EntityType{}.Mixin()
	entitytypeMixinHooks0 := entitytypeMixin[0].Hooks()
	entitytype.Hooks[0] = entitytypeMixinHooks0[0]
	entitytypeMixinInters0 := entitytypeMixin[0].Interceptors()
	entitytype.Interceptors[0] = entitytypeMixinInters0[0]
	entitytypeMixinFields0 := entitytypeMixin[0].Fields()
	_ = entitytypeMixinFields0
	entitytypeFields := schema.EntityType{}.Fields()
	_ = entitytypeFields
	// entitytypeDescCreatedAt is the schema descriptor for created_at field.
	entitytypeDescCreatedAt := entitytypeMixinFields0[0].Descriptor()
	// entitytype.DefaultCreatedAt holds the default value on creation for the created_at field.
	entitytype.DefaultCreatedAt = entitytypeDescCreatedAt.Default.(func() time.Time)
	// entitytypeDescUpdatedAt is the schema descriptor for updated_at field.
	entitytypeDescUpdatedAt := entitytypeMixinFields0[1].Descriptor()
	// entitytype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	entitytype.DefaultUpdatedAt = entitytypeDescUpdatedAt.Default.(func() time.Time)
	// entitytype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	entitytype.UpdateDefaultUpdatedAt = entitytypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	friendshipMixin := schema.Friendship{}.Mixin()
	friendshipMixinHooks0 := friendshipMixin[0].Hooks()
	friendship.Hooks[0] = friendshipMixinHooks0[0]
	friendshipMixinInters0 := friendshipMixin[0].Interceptors()
	friendship.Interceptors[0] = friendshipMixinInters0[0]
	friendshipMixinFields0 := friendshipMixin[0].Fields()
	_ = friendshipMixinFields0
	friendshipFields := schema.Friendship{}.Fields()
	_ = friendshipFields
	// friendshipDescCreatedAt is the schema descriptor for created_at field.
	friendshipDescCreatedAt := friendshipMixinFields0[0].Descriptor()
	// friendship.DefaultCreatedAt holds the default value on creation for the created_at field.
	friendship.DefaultCreatedAt = friendshipDescCreatedAt.Default.(func() time.Time)
	// friendshipDescUpdatedAt is the schema descriptor for updated_at field.
	friendshipDescUpdatedAt := friendshipMixinFields0[1].Descriptor()
	// friendship.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	friendship.DefaultUpdatedAt = friendshipDescUpdatedAt.Default.(func() time.Time)
	// friendship.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	friendship.UpdateDefaultUpdatedAt = friendshipDescUpdatedAt.UpdateDefault.(func() time.Time)
	friendshipstatusMixin := schema.FriendshipStatus{}.Mixin()
	friendshipstatusMixinHooks0 := friendshipstatusMixin[0].Hooks()
	friendshipstatus.Hooks[0] = friendshipstatusMixinHooks0[0]
	friendshipstatusMixinInters0 := friendshipstatusMixin[0].Interceptors()
	friendshipstatus.Interceptors[0] = friendshipstatusMixinInters0[0]
	friendshipstatusMixinFields0 := friendshipstatusMixin[0].Fields()
	_ = friendshipstatusMixinFields0
	friendshipstatusFields := schema.FriendshipStatus{}.Fields()
	_ = friendshipstatusFields
	// friendshipstatusDescCreatedAt is the schema descriptor for created_at field.
	friendshipstatusDescCreatedAt := friendshipstatusMixinFields0[0].Descriptor()
	// friendshipstatus.DefaultCreatedAt holds the default value on creation for the created_at field.
	friendshipstatus.DefaultCreatedAt = friendshipstatusDescCreatedAt.Default.(func() time.Time)
	// friendshipstatusDescUpdatedAt is the schema descriptor for updated_at field.
	friendshipstatusDescUpdatedAt := friendshipstatusMixinFields0[1].Descriptor()
	// friendshipstatus.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	friendshipstatus.DefaultUpdatedAt = friendshipstatusDescUpdatedAt.Default.(func() time.Time)
	// friendshipstatus.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	friendshipstatus.UpdateDefaultUpdatedAt = friendshipstatusDescUpdatedAt.UpdateDefault.(func() time.Time)
	notificationMixin := schema.Notification{}.Mixin()
	notificationMixinHooks0 := notificationMixin[0].Hooks()
	notification.Hooks[0] = notificationMixinHooks0[0]
	notificationMixinInters0 := notificationMixin[0].Interceptors()
	notification.Interceptors[0] = notificationMixinInters0[0]
	notificationMixinFields0 := notificationMixin[0].Fields()
	_ = notificationMixinFields0
	notificationFields := schema.Notification{}.Fields()
	_ = notificationFields
	// notificationDescCreatedAt is the schema descriptor for created_at field.
	notificationDescCreatedAt := notificationMixinFields0[0].Descriptor()
	// notification.DefaultCreatedAt holds the default value on creation for the created_at field.
	notification.DefaultCreatedAt = notificationDescCreatedAt.Default.(func() time.Time)
	// notificationDescUpdatedAt is the schema descriptor for updated_at field.
	notificationDescUpdatedAt := notificationMixinFields0[1].Descriptor()
	// notification.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notification.DefaultUpdatedAt = notificationDescUpdatedAt.Default.(func() time.Time)
	// notification.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notification.UpdateDefaultUpdatedAt = notificationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notificationDescIsRead is the schema descriptor for isRead field.
	notificationDescIsRead := notificationFields[0].Descriptor()
	// notification.DefaultIsRead holds the default value on creation for the isRead field.
	notification.DefaultIsRead = notificationDescIsRead.Default.(bool)
	notificationchangeMixin := schema.NotificationChange{}.Mixin()
	notificationchangeMixinHooks0 := notificationchangeMixin[0].Hooks()
	notificationchange.Hooks[0] = notificationchangeMixinHooks0[0]
	notificationchangeMixinInters0 := notificationchangeMixin[0].Interceptors()
	notificationchange.Interceptors[0] = notificationchangeMixinInters0[0]
	notificationchangeMixinFields0 := notificationchangeMixin[0].Fields()
	_ = notificationchangeMixinFields0
	notificationchangeFields := schema.NotificationChange{}.Fields()
	_ = notificationchangeFields
	// notificationchangeDescCreatedAt is the schema descriptor for created_at field.
	notificationchangeDescCreatedAt := notificationchangeMixinFields0[0].Descriptor()
	// notificationchange.DefaultCreatedAt holds the default value on creation for the created_at field.
	notificationchange.DefaultCreatedAt = notificationchangeDescCreatedAt.Default.(func() time.Time)
	// notificationchangeDescUpdatedAt is the schema descriptor for updated_at field.
	notificationchangeDescUpdatedAt := notificationchangeMixinFields0[1].Descriptor()
	// notificationchange.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notificationchange.DefaultUpdatedAt = notificationchangeDescUpdatedAt.Default.(func() time.Time)
	// notificationchange.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notificationchange.UpdateDefaultUpdatedAt = notificationchangeDescUpdatedAt.UpdateDefault.(func() time.Time)
	notificationobjectMixin := schema.NotificationObject{}.Mixin()
	notificationobjectMixinHooks0 := notificationobjectMixin[0].Hooks()
	notificationobject.Hooks[0] = notificationobjectMixinHooks0[0]
	notificationobjectMixinInters0 := notificationobjectMixin[0].Interceptors()
	notificationobject.Interceptors[0] = notificationobjectMixinInters0[0]
	notificationobjectMixinFields0 := notificationobjectMixin[0].Fields()
	_ = notificationobjectMixinFields0
	notificationobjectFields := schema.NotificationObject{}.Fields()
	_ = notificationobjectFields
	// notificationobjectDescCreatedAt is the schema descriptor for created_at field.
	notificationobjectDescCreatedAt := notificationobjectMixinFields0[0].Descriptor()
	// notificationobject.DefaultCreatedAt holds the default value on creation for the created_at field.
	notificationobject.DefaultCreatedAt = notificationobjectDescCreatedAt.Default.(func() time.Time)
	// notificationobjectDescUpdatedAt is the schema descriptor for updated_at field.
	notificationobjectDescUpdatedAt := notificationobjectMixinFields0[1].Descriptor()
	// notificationobject.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notificationobject.DefaultUpdatedAt = notificationobjectDescUpdatedAt.Default.(func() time.Time)
	// notificationobject.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notificationobject.UpdateDefaultUpdatedAt = notificationobjectDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notificationobjectDescEntityID is the schema descriptor for entity_id field.
	notificationobjectDescEntityID := notificationobjectFields[0].Descriptor()
	// notificationobject.EntityIDValidator is a validator for the "entity_id" field. It is called by the builders before save.
	notificationobject.EntityIDValidator = notificationobjectDescEntityID.Validators[0].(func(int) error)
	userMixin := schema.User{}.Mixin()
	userMixinHooks0 := userMixin[0].Hooks()
	user.Hooks[0] = userMixinHooks0[0]
	userMixinInters0 := userMixin[0].Interceptors()
	user.Interceptors[0] = userMixinInters0[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)
