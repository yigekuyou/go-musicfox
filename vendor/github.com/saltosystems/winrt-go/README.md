# The WinRT Go language projection

WinRT Go is an autogenerated language projection for Windows Runtime (WinRT) APIs.

The project was originally intended to assist in the development of the Windows implementation of the [tinygo-org/bluetooth](https://github.com/tinygo-org/bluetooth) library.
That's why it currently only contains WinRT APIs related to Bluetooth, but we are open to include new APIs if requested.

> [!IMPORTANT]
> Due to the nature of this project we decided not to tag any releases.
> All commits to the `main` branch are considered stable, and we encourage you to use the latest commit available.
> 
> Breaking changes in the `main` branch should be an exception, but we may include them in favor of the simplicity of the project.

The project also contains the `winrt-go-gen` code generator and the source Windows Metadata (WinMD) files (located in `/internal/winmd/metadata`) that describe the WinRT APIs.

Notice that the code generator is not capable of generating any WinRT types.
It will work with most of them, but we've added the functionalities we exclusively needed to generate the BLE APIs.
So some things may still be missing.
Check out the [known missing features](#known-missing-features) section below for more information.

## Generated code

All the generated code is stored in the `windows` folder, and is divided in folders that match the namespace of each class.

Since our goal was to only use a subset of the WinRT APIs, the generated classes may only contain a subset of their methods.
This helps us generate less code, because we can remove dependencies between classes and completely skip their generation.

Also notice that the generated method names may differ from the ones defined in the WinRT API.
This is because Go does not support method overloading, so we are using the overload name defined in the WinMD files.
The [`GetGattServicesAsync` method](https://docs.microsoft.com/en-us/uwp/api/windows.devices.bluetooth.bluetoothledevice.getgattservicesasync?view=winrt-22621), for example, has an attribute defining the following overload name: `GetGattServicesWithCacheModeAsync`.

This also affects static methods, which include their class name as prefix to avoid collisions between classes inside the same package.

## Generating the code

The code is generated using `go generate`. But the Makefile includes a target (`make gen-files`) that removes all generated code and executes the `go generate` command.

You can also call the code generator manually.

```
Usage of winrt-go-gen:
  -class string
        The class to generate. This should include the namespace and the class name, e.g. 'System.Runtime.InteropServices.WindowsRuntime.EventRegistrationToken'.
  -config string
        config file (optional)
  -debug
        Enables the debug logging.
  -method-filter value
        The filter to use when generating the methods. This option can be set several times, 
        the given filters will be applied in order, and the first that matches will determine the result. The generator
        will allow any method by default. The filter uses the overloaded method name to discriminate between overloaded
        methods.
    
        You can use the '!' character to negate a filter. For example, to generate all methods except the 'Add' method:
            -method-filter !Add
    
        You can also use the '*' character to match any method, so if you want to generate only the 'Add' method, you can do:
            -method-filter Add -method-filter !*
```

## Known missing features

- If an interface extends another one, the methods of the parent interface are not generated.
- There are still some unsupported data types:
    - Multi-dimensional arrays (`ELEMENT_TYPE_ARRAY`)
    - References (`ELEMENT_TYPE_BYREF`)
    - Pointer to functions (`ELEMENT_TYPE_FNPTR`)
    - Pointer types (`ELEMENT_TYPE_PTR`)
    - Typed references (`ELEMENT_TYPE_TYPEDBYREF`)