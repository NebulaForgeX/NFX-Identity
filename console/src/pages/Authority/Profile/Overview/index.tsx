import { ProfileKindEnum } from "nfx-ui/enums";

import { OverviewView } from "@/components/identity";
import { scopePaths } from "@/navigations";

export default function AuthorityProfileOverviewPage() {
  return <OverviewView paths={scopePaths(ProfileKindEnum.AUTHORITY)} />;
}
