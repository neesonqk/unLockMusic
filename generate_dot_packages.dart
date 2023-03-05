import 'dart:convert';
import 'dart:io';

Future<void> main() async {
  File file = File(".dart_tool/package_config.json");
  dynamic pkgconfig = jsonDecode(await file.readAsString());
  List<dynamic> pkgs = pkgconfig["packages"];
  for (var elm in pkgs) {
    if (elm["rootUri"] == "../") {
      elm["rootUri"] = "./";
    }
  }
  pkgconfig["packages"] = pkgs;
  print(jsonEncode(pkgconfig));
}
