"""This module provides a rule to make a file executable."""

def _make_executable_impl(ctx):
    # Define the output file as a new executable file
    output = ctx.actions.declare_file(ctx.attr.src.basename)

    # Run a shell command to make the file executable
    ctx.actions.run_shell(
        inputs = [ctx.file.src],
        outputs = [output],
        command = "cp {input} {output} && chmod +x {output}".format(
            input = ctx.file.src.path,
            output = output.path,
        ),
    )

    # Return the new executable file as the output
    return [DefaultInfo(files = depset([output]), executable = output)]

make_executable = rule(
    implementation = _make_executable_impl,
    attrs = {
        "src": attr.label(allow_single_file = True),
    },
)
