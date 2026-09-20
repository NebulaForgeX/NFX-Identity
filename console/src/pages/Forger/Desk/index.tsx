import { ProfileKindEnum } from "nfx-ui/enums";

import { DeskView } from "@/components/identity";
import { scopePaths } from "@/navigations";

export default function ForgerDeskPage() {
  return <DeskView paths={scopePaths(ProfileKindEnum.FORGER)} />;
}
