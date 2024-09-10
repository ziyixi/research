"""This module provides a rule to make a file executable."""

def _make_executable_impl(ctx):
    # Use the original file and make it executable
    ctx.actions.run_shell(
        inputs = [ctx.file.src],
        outputs = [ctx.actions.declare_file(ctx.file.src.basename)],
        command = "chmod +x {input}".format(input = ctx.file.src.path),
    )

    # Return the original file as the executable output
    return [DefaultInfo(executable = ctx.file.src)]

make_executable = rule(
    implementation = _make_executable_impl,
    attrs = {
        "src": attr.label(allow_single_file = True),
    },
)
