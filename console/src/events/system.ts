import { EventEmitter, defineEvents, type EventNamesOf } from "nfx-ui/events";
import { singleton } from "nfx-ui/utils";

export const systemEvents = defineEvents({
  INVALIDATE_SYSTEM_STATE: "SYSTEM:INVALIDATE_SYSTEM_STATE",
  INVALIDATE_SYSTEM_STATES: "SYSTEM:INVALIDATE_SYSTEM_STATES",
});

type SystemEvent = EventNamesOf<typeof systemEvents>;

class SystemEventEmitter extends EventEmitter<SystemEvent> {
  constructor() {
    super(systemEvents);
  }
}

export const systemEventEmitter = new (singleton(SystemEventEmitter))();
